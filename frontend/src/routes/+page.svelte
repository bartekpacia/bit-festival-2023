<script lang="ts">
	let mediaRecorder: MediaRecorder | null = null;
	let isRecording = false;
	let chunks: Blob[] = [];
	let downloadHref = "";

	if (typeof window !== "undefined") {
		navigator.mediaDevices.getUserMedia({ audio: true }).then((stream) => {
			mediaRecorder = new MediaRecorder(stream);

			mediaRecorder.ondataavailable = (e) => {
				chunks.push(e.data);
			};

			mediaRecorder.onstop = (e) => {
				console.log(
					"data available after MediaRecorder.stop() called.",
				);

				const clipName = "test";

				const blob = new Blob(chunks, {
					type: "audio/ogg; codecs=opus",
				});

				chunks = [];
				downloadHref = URL.createObjectURL(blob);
			};
		});
	}

	function startStop() {
		if (!mediaRecorder) {
			return;
		}

		if (isRecording) {
			mediaRecorder.stop();
			isRecording = false;
		} else {
			console.log("Start");
			mediaRecorder.start();
			isRecording = true;
		}
	}
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<h1>Hello world</h1>
<button on:click={startStop}>
	{#if isRecording}
		Stop
	{:else}
		Start
	{/if}
</button>

<a download="cos.ogg" href={downloadHref}>Download</a>
