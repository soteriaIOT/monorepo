import sys
sys.path.append("./common")
sys.path.append("..")
from common.kafka_helper import produce_message
from common.kafka_helper import consume_confluence
from common.kafka_helper import read_requirements
from common.tags import COMMON_TAGS


class Patcher:
    def _update_requirements(self, updated_requirements: str):
        """
        Update the requirements file.
        """
        with open("../app/requirements.txt", "w") as f:
            f.write(updated_requirements)
    
    def _rebuild_service(self):
        """
        Rebuild the service.
        """
        print("REBUILD")

    def _stop_service(self):
        """
        Stop the service.
        """
        print("STOP")

        
    def _start_service(self):
        """
        Restart the service.
        """
        print("START")
    
    def watch(self):
        """
        Watch for changes to the requirements file.
        """
        DEVICE_ID = COMMON_TAGS.get("host")
        
        REQUIREMENTS = read_requirements("../app/requirements.txt")
        produce_message("requirements", DEVICE_ID, "\n".join(REQUIREMENTS)) 
        for message in consume_confluence("device-updates"):
            key = message.key()
            value = message.value()
            if key is None:
                print("Key: {}".format(key))
                continue
            if value is None:
                print("Value: {}".format(value))
                continue

            device, package_and_version = key.decode("utf-8"), value.decode("utf-8") 
            if str(device) == DEVICE_ID:
                print("Got update for device:", device, package_and_version)
                package, version = package_and_version.split("==")
                updated_requirements = "\n".join([
                    line if not line.startswith(package)
                    else "{}=={}".format(package, version) for line in REQUIREMENTS
                    
                ])
                self._update_requirements(updated_requirements)
                self._stop_service()
                self._rebuild_service()
                self._start_service()
                REQUIREMENTS = read_requirements("../app/requirements.txt")
                produce_message("requirements", DEVICE_ID, "\n".join(REQUIREMENTS)) 

            print(device, package_and_version)
        
        




if __name__ == "__main__":
    patcher = Patcher()
    patcher.watch()
