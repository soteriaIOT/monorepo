
import subprocess
import os 
import signal

PATH = "~/github/fydp/monorepo/"

def update_requirements(updated_requirements: str):
    """
    Update the requirements file.
    """
    with open("requirements.txt", "w") as f:
        f.write(updated_requirements)

def run_pip_install():
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

def stop_service():
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

def start_service():
    """
    Restart the service.
    """
    command = f"""cd {PATH}/device/app;
        source bin/activate;
        python3 main.py &
    """
    ret = subprocess.Popen(command)

if __name__ == "__main__":
    stop_service()
    updated_requirements = """backoff==1.11.1
certifi==2021.10.8
chardet==3.0.4
idna==2.7
psutil==5.9.0
python-dotenv==0.19.2
ratelimit==2.2.1
requests==2.20.0
speedtest-cli==2.1.3
urllib3==1.24.3"""
    update_requirements(updated_requirements)
    run_pip_install()
    start_service()
    
