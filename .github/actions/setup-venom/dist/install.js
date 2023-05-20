"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.Install = void 0;
const core = __importStar(require("@actions/core"));
const exec = __importStar(require("@actions/exec"));
class Install {
    install(venomVersion, venomWebVersion, plateform) {
        return new Promise((resolve, reject) => __awaiter(this, void 0, void 0, function* () {
            core.info(`Install venom version ${venomVersion} with venomWeb version ${venomWebVersion}`);
            try {
                // Checkout venom
                core.info(`Clone venom repository`);
                yield exec.exec("git", ["clone", "-b", venomVersion, "https://github.com/ovh/venom.git", "venom"]);
                // Get venom web specific version
                core.info(`Get venomWeb dependency`);
                let options = {};
                options.cwd = "./venom";
                const dependency = "github.com/kevinramage/venomWeb@" + venomWebVersion;
                core.info(dependency);
                yield exec.exec("go", ["get", dependency], options);
                // Checkout venom web
                //await exec.exec("git", ["clone", "-b", venomWebVersion, "https://github.com/kevinramage/venomWeb.git", "venomWeb"]);
                // Compile venom for target plateform (windows-latest, ubuntu-latest, macos-latest)
                core.info(`Compile`);
                options.cwd = "./venom/cmd/venom";
                let targetPlateform = "", execName = "";
                if (plateform == "ubuntu-latest") {
                    targetPlateform = "OS=linux";
                    execName = "venom.linux-amd64";
                }
                else if (plateform == "macos-latest") {
                    targetPlateform = "OS=darwin";
                    execName = "venom.darwin-amd64";
                }
                else if (plateform == "windows-latest") {
                    targetPlateform = "OS=windows";
                    execName = "venom.windows-amd64";
                }
                yield exec.exec("make", ["build", targetPlateform, "ARCH=amd64"], options);
                // Rename build
                yield exec.exec("mv", ["dist/" + execName, "venom"], options);
                resolve();
            }
            catch (err) {
                core.error(`An error during installation of venom: ${err}`);
                reject(err);
            }
        }));
    }
}
exports.Install = Install;
