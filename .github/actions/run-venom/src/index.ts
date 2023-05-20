import * as core from "@actions/core";
import * as exec from "@actions/exec";
import { Plateform, SYSTEM_TYPE } from "./plateform";


class Index {

    async run() {
        return new Promise<void>(async (resolve, reject) => {
            try {
    
                // Detect plateform
                core.info("Detect plateform")
                const plateform = new Plateform();
                plateform.detectPlateform();

                // Copy prerequisites
                core.info("Copy prerequisites")
                if (plateform.getSystem() == SYSTEM_TYPE.WINDOWS) {

                } else {
                    await exec.exec("mv venom/cmd/venom/venom venomWeb/venom");
                    await exec.exec("cp /opt/chromedriver/chromedriver ./venomWeb/chromedriver")
                }

                // Run venom
                core.info("Run venom");
                let options : any = {};
                options.cwd = "venomWeb";
                let cmdLine = "";
                if (plateform.getSystem() == SYSTEM_TYPE.WINDOWS) {
                    cmdLine = ".\\venom -vvv --format=xml run tests\\windows\\chrome\\core.yml";
                } else if (plateform.getSystem() == SYSTEM_TYPE.LINUX) {
                    cmdLine = "./venom -vvv --format=xml run tests/linux/chrome/core.yml";
                } else if (plateform.getSystem() == SYSTEM_TYPE.DARWIN) {
                    cmdLine = "./venom -vvv --format=xml run tests/darwin/chrome/core.yml";
                } else {
                    throw "Invalid plateform: " + plateform.getSystem();
                }

                resolve();
                
            } catch (err) {
                core.setFailed(String(err));
                reject(err);
            }
        })
    }
}

new Index().run();