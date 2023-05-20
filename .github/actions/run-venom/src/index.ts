import * as core from "@actions/core";
import * as exec from "@actions/exec";
import * as fs from "fs";
import { Plateform, SYSTEM_TYPE } from "./plateform";
import { Buffer } from "buffer";


class Index {

    async run() {
        return new Promise<void>(async (resolve, reject) => {
            try {
                // Options
                let output = "";
                let options : any = {};
                options.listeners = {
                    stdout: (data: Buffer) => {
                        output += data.toString();
                    },
                };
    
                // Detect plateform
                core.info("Detect plateform")
                const plateform = new Plateform();
                plateform.detectPlateform();

                // Temp
                output = "";
                if (plateform.getSystem() == SYSTEM_TYPE.WINDOWS) {
                    await exec.exec("dir .\\venom\\cmd\\venom\\venom", [], options);
                    await exec.exec("dir .\\venom\\cmd\\venom\\dist", [], options);
                    await exec.exec("dir \"C:\\Program\ Files\\chromedriver\"", [], options);
                } else {
                    await exec.exec("ls -la .", [], options);
                }
                core.info("Output:");
                core.info(output);

                // Copy prerequisites
                core.info("Copy prerequisites")
                if (plateform.getSystem() == SYSTEM_TYPE.WINDOWS) {
                    core.info("Copy venom");
                    await fs.promises.copyFile(".\\venom\\cmd\\venom\\venom", ".\\venomWeb\\");
                    core.info("Copy venom driver");
                    await fs.promises.copyFile("\"C:\\Program\ Files\\chromedriver\\chromedriver\"", ".\\venomWeb\\");
                } else {
                    await exec.exec("mv venom/cmd/venom/venom venomWeb/venom");
                    await exec.exec("cp /opt/chromedriver/chromedriver ./venomWeb/chromedriver")
                }

                // Run venom
                core.info("Run venom");
                options.cwd = "venomWeb";

                // Temp
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