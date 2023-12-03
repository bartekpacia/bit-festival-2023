import express from "express";
import { FileSource, createClient } from "@deepgram/sdk";
import dotenv from "dotenv";
import cors from "cors";
import fs from "fs"
import multer from "multer";
import http from "http"

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
    fetch("http://127.0.0.1:11434/api/generate", {
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

app.post("/speach-to-json", upload.single('file'), async (req, res) => {
    const { file } = req

    console.log(file)
    try {
        if (file && file.buffer) {
            const prompt = await transcriber(file.buffer)
            const system = fs.readFileSync("./jsonGenerationInput.txt", 'utf-8');
            console.log("transcribed prompt: " + prompt)

            const requestBodyData = JSON.stringify({
                model: "llama2",
                format: "json",
                stream: false,
                prompt,
                system
            })

            const options: http.RequestOptions = {
                hostname: '127.0.0.1',
                port: 11434,
                timeout: 1000 * 60 * 60,
                path: '/api/generate',
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Content-Length': Buffer.byteLength(requestBodyData)
                }
            };

            const req = http.request(options, resp => {
                let data = "";

                // A chunk of data has been recieved.
                resp.on("data", chunk => {
                    data += chunk;
                });

                // The whole response has been received. Print out the result.
                resp.on("end", () => {
                    const lammaJson = JSON.parse(data);
                    console.log(lammaJson.response)
                    res.json(JSON.parse(lammaJson.response))
                });
            })
                .on("error", err => {
                    console.log("Error: " + err.message);
                });

            // Wysłanie danych w ciele żądania
            req.write(requestBodyData);

            // Zakończenie żądania
            req.end();

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
