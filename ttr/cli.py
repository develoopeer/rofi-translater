import subprocess
import typer

app = typer.Typer()

@app.command()
def run(name: str):
    try:
        proc = subprocess.Popen(["python3 -m ttr.main > /dev/null"] , shell=True)
        print("Tray translater is now running: {0}".format(proc.pid))
    except Exception:
        print("Another instance of tray translater is running")

if __name__ == "__main__":
    app()