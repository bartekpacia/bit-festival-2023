import express from "express";
import { FileSource, createClient } from "@deepgram/sdk";
import dotenv from "dotenv";
import cors from "cors";
import fs from "fs"
import multer from "multer";

const app = express();
app.use(express.json({ limit: "1024mb" }));
dotenv.config();
const upload = multer()

const deepgram = createClient(process.env.DEEPRAM_API || "");

const transcriber = async (file: FileSource) => new Promise<string>(async (res, rej) => {
    try {
        const { result, error } = await deepgram.listen.prerecorded.transcribeFile(
            file,
            {
                model: "enhanced",
                language: "pl"
            }
        );
        if (result) {
            res(result.results.channels[0].alternatives[0].transcript)
        }
        rej(error)
    } catch (err) {
        rej(err)
    }
})

app.use(cors());

app.get("/", (req, res) => {
    res.send("works")
});
app.get("/testfile", (req, res) => {
    const fileSrc = fs.readFileSync("./test.m4a");
    transcriber(fileSrc)
        .then(o => {
            console.log(o)
            res.json(o)
        })
        .catch(err => {
            console.error(err)
            res.json(err)
        })
});

app.get("/test-lamma", (req, res) => {
    fetch("http://localhost:11434/api/generate", {
        method: "POST",
        body: JSON.stringify({
            "model": "llama2",
            "prompt": "give me example of not empty json",
            "format": "json",
            "stream": false
        })
    })
        .then(async (o) => {
            console.log("somehow it is okay")
            const a = await o.json()
            console.log(o, "-----\n\n----\n\n", a)
            res.json(a)
        })
        .catch(err => {
            console.log("error accoured")
            console.error(err)
            res.send(err)
        })
})

app.get("/speach-to-json", upload.single('file'), async (req, res) => {
    const { file } = req

    console.log(file)
    try {
        if (file && file.buffer) {
            const transText = await transcriber(file.buffer)

            let prompt = fs.readFileSync("./jsonGenerationInput.txt", 'utf-8');
            prompt = prompt.replace("%generated_text%", transText)

            console.log("transcribed prompt:\n\n", prompt)
            const lammaAns = await fetch("http://localhost:11434/api/generate", {
                method: "POST",
                body: JSON.stringify({
                    "model": "llama2",
                    "prompt": prompt,
                    "format": "json",
                    "stream": false
                })
            })
            const lammaJson = await lammaAns.json()
            console.log(lammaJson.response)
            res.json(JSON.parse(lammaJson.response))
        } else {
            res.status(402).send("no file provided")
        }
    } catch (err) {
        res.status(400).send(err)
    }
})

app.get("/voice-file-to-text", upload.single('file'), (req, res) => {
    const { file } = req
    console.log(file)
    if (file && file.buffer) {
        transcriber(file.buffer)
            .then(o => {
                console.log(o)
                res.json(o)
            })
            .catch(err => {
                console.error(err)
                res.json(err)
            })
    }
});

app.listen(9999, () => console.log(`listening on port 9999`));
