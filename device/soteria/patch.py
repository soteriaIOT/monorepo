from os import path

from docker import from_env
from docker import errors

from configparser import ConfigParser
from soteria.kafka_helper import produce_message
from soteria.kafka_helper import consume_confluence
from common.tags import COMMON_TAGS

DEVICE_UPDATES = "device-updates"
DEVICE_REQUIREMENTS = "device-requirements"
DEVICE_ID = COMMON_TAGS.get("host")

class Patcher:
    def __init__(self):
        self.config = ConfigParser()
        self.config.read("./soteria.ini")
        self.config = self.config["DEFAULT"]
        self.app_path = self.config.get("PathToApp")
        self.dependency_file = path.join(self.app_path, self.config.get("PathToRequirements"))
        self.docker_file_folder = path.dirname(path.join(self.app_path, self.config.get("PathToDockerfile")))
        self.__init__docker()
    
    def __init__docker(self):
        """
        Initialize the docker client.
        """
        self.docker = from_env()
        try:
            self.image = self.docker.images.get("device-app:latest")
        except errors.ImageNotFound:
            self._build_service()
        try:
            self.container = self.docker.containers.get("device-app")
        except errors.NotFound:
            self._start_service()
        self.docker_init_done = True

    def _update_requirements(self, updated_requirements: str):
        """
        Update the requirements file.
        """
        with open(self.dependency_file, "w") as f:
            f.write(updated_requirements)

    def _build_service(self):
        """
        Build the service.
        """
        self.image, _ = self.docker.images.build(
            path=self.docker_file_folder,
            network_mode="host",
            tag="device-app:latest"
        )

    def _stop_service(self):
        """
        Stop the service.
        """
        self.container.stop()
        self.docker.containers.prune()

    def _start_service(self):
        """
        Restart the service.
        """
        self.container = self.docker.containers.run(self.image, detach=True, auto_remove=True, name="device-app")
    
    def _read_requirements(self):
        """
        Read the requirements file.
        """
        return self.container.exec_run(
            cmd="pip freeze", 
            stream=False, 
            stdout=True, 
            stdin=False
        ).output.decode("utf-8").split("\n")

    def update(self, updated_requirements):
        print("Updating requirements...")
        self._stop_service()
        self._update_requirements(updated_requirements=updated_requirements)
        self._build_service()
        self._start_service()
        print("Done updating dependencies")

    def watch(self):
        """
        Watch for changes to the requirements file.
        """
        
        REQUIREMENTS = self._read_requirements()
        produce_message(DEVICE_REQUIREMENTS, DEVICE_ID, REQUIREMENTS)
        for message in consume_confluence(DEVICE_UPDATES):
            key = message.key()
            value = message.value()
            try:
                device, package_and_version = key.decode("utf-8"), value.decode("utf-8")
            except Exception as e:
                print("Exception: {}".format(e))
                continue
            if str(device) == DEVICE_ID:
                print("Got update for device:", device, package_and_version)
                package, version = package_and_version.split("==")
                updated_requirements = "\n".join(
                    [
                        line
                        if not line.split("==")[0] == package
                        else "{}=={}".format(package, version)
                        for line in REQUIREMENTS
                    ]
                )
                self.update(updated_requirements)
                REQUIREMENTS = self._read_requirements()
                produce_message(
                    DEVICE_REQUIREMENTS, DEVICE_ID, REQUIREMENTS
                )


if __name__ == "__main__":
    patcher = Patcher()
    patcher.watch()
