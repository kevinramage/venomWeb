import * as core from "@actions/core";
import * as exec from "@actions/exec";
import { Plateform, SYSTEM_TYPE } from "./plateform";
import { Buffer } from "buffer";


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
                    await exec.exec("copy venom\\cmd\\venom\\venom venomWeb\\venom");
                    await exec.exec("cp \"C:\\Program\ Files\\chromedriver\\chromedriver\" venomWeb\\chromedriver")
                } else {
                    await exec.exec("mv venom/cmd/venom/venom venomWeb/venom");
                    await exec.exec("cp /opt/chromedriver/chromedriver ./venomWeb/chromedriver")
                }

                // Run venom
                core.info("Run venom");
                let output = "";
                let options : any = {};
                options.listeners = {
                    stdout: (data: Buffer) => {
                        output += data.toString();
                    },
                };
                options.cwd = "venomWeb";

                output = "";
                if (plateform.getSystem() == SYSTEM_TYPE.WINDOWS) {
                    await exec.exec("dir .", [], options);
                } else {
                    await exec.exec("ls -la .", [], options);
                }
                core.info("Output:");
                core.info(output);

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
                core.info(cmdLine);
                await exec.exec(cmdLine, [], options);
                core.info("Output:");
                core.info(output);
                
                // Get result
                output = "";
                if (plateform.getSystem() == SYSTEM_TYPE.WINDOWS) {
                    await exec.exec("dir .", [], options);
                } else {
                    await exec.exec("ls -la .", [], options);
                }
                core.info("Output:");
                core.info(output);


                resolve();
                
            } catch (err) {
                core.setFailed(String(err));
                reject(err);
            }
        })
    }
}

new Index().run();