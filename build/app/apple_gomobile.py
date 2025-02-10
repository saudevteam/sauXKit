import os
import subprocess

from app.build import Builder


class AppleGoMobileBuilder(Builder):
    def before_build(self):
        super().before_build()
        self.clean_lib_dirs(["SauXKit.xcframework"])
        self.prepare_gomobile()

    def build(self):
        self.before_build()

        os.chdir(self.lib_dir)
        ret = subprocess.run(
            [
                "gomobile",
                "bind",
                "-target",
                "ios,iossimulator,macos",
                "-iosversion",
                "15.0",
            ]
        )
        if ret.returncode != 0:
            raise Exception("build failed")
