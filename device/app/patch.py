
import subprocess
import os 
import signal

from http.server import HTTPServer
from http.server import BaseHTTPRequestHandler
from urllib.parse import urlparse
from urllib.parse import parse_qs

PATH = "~/soteria/monorepo/"

class PatchServer(BaseHTTPRequestHandler):

    def _update_requirements(self, updated_requirements: str):
        """
        Update the requirements file.
        """
        with open("requirements.txt", "w") as f:
            f.write(updated_requirements)

    def _run_pip_install(self):
        """
        Run pip install.
        """
        command = f"""
            cd {PATH}device/app;
            source bin/activate;
            pip3 install -r requirements.txt;
        """
        ret = subprocess.run(command, capture_output=True, shell=True)
        print(ret.stdout.decode())

    def _stop_service(self):
        """
        Stop the service.
        """
        
        def get_pid(name):
            # https://stackoverflow.com/questions/26688936/how-to-get-pid-by-process-name
            for line in os.popen("ps aux | grep 'python3 main.py' | grep -v grep"):
                pid = int(line.split()[1])
                return pid
        
        def kill_pid(pid: int):
            # https://stackoverflow.com/questions/12309269/how-to-kill-a-process-started-with-subprocess-in-python
            print("Killing process with pid {}".format(pid))
            os.kill(pid, signal.SIGKILL)
        
        pid = get_pid("python3")
        kill_pid(pid)

    def _start_service(self):
        """
        Restart the service.
        """
        command = f"""sh spawn.sh"""
        ret = subprocess.run(command, shell=True)
    

    def do_GET(self):
        query = parse_qs(urlparse(self.path).query)
        updated_requirements = "\n".join([f"{key}=={query[key][0]}" for key in query])
        self._stop_service()

        self._update_requirements(updated_requirements)
        self._run_pip_install()
        self._start_service()
        self.send_response(200)
        self.send_header('Content-type', 'text/html')
        self.end_headers()
        self.wfile.write(bytes(updated_requirements, "utf8"))
        return

server = HTTPServer(("localhost", 7000), PatchServer)
server.serve_forever()



if __name__ == "__main__":
    P = PatchServer()
    P._stop_service()
    updated_requirements = """backoff==1.11.1
certifi==2021.10.8
chardet==3.0.4
idna==2.7
psutil==5.9.0
python-dotenv==0.19.2
ratelimit==2.2.1
requests==2.20.0
speedtest-cli==2.1.3
urllib3==1.26.5"""
    P._update_requirements(updated_requirements)
    P._run_pip_install()
    P._start_service()
    
