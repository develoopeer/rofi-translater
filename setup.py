from setuptools import setup
from setuptools.command.develop import develop
from setuptools.command.install import install
from subprocess import check_call


class PostDevelopCommand(develop):
    """Post-installation for development mode."""
    def run(self):
        develop.run(self)
        check_call("./install.sh".split())

class PostInstallCommand(install):
    """Post-installation for installation mode."""
    def run(self):
        install.run(self)
        check_call("apt-get install this-package".split())

setup(
    name="ttr",
    version="0.0.1",
    packages=["ttr"],
    package_data={
        "ttr": ["*.wav" , "*.png" , "*.service" , "*.sh"]
    },
    include_package_data=True
)