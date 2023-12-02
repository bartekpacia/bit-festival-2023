<script lang="ts">
	import mic from "$lib/icons/mic.svg";
	import calculate from "$lib/icons/calculate.svg";

	enum AssistantState {
		ReadyForInput,
		Listening,
		Processing,
	}

	enum AmpacityOrMaxPower {
		Ampacity,
		MaxPower,
	}

	enum Placements {
		A1,
		A2,
		B2,
		C,
		E,
	}

	let mediaRecorder: MediaRecorder | null = null;
	let chunks: Blob[] = [];
	let downloadHref = "";
	let assistantState: AssistantState = AssistantState.ReadyForInput;
	let ampacityOrMaxPower = AmpacityOrMaxPower.Ampacity;
	let assistantBtnColor: string;

	let ampacity = 0;
	let maxPower = 0;
	let veinsUnderLoad = 0;
	let placements: Placements = Placements.A1;
	let temperature = 0;

	$: if (assistantState === AssistantState.ReadyForInput) {
		assistantBtnColor = "rgb(0, 33, 95)";
	} else if (assistantState === AssistantState.Listening) {
		assistantBtnColor = "green";
	} else {
		assistantBtnColor = "rgb(255, 92, 0)";
	}

	if (typeof window !== "undefined") {
		navigator.mediaDevices.getUserMedia({ audio: true }).then((stream) => {
			mediaRecorder = new MediaRecorder(stream);

			mediaRecorder.onstart = (e) => {
				console.log("Start");
				assistantState = AssistantState.Listening;
			};

			mediaRecorder.ondataavailable = (e) => {
				chunks.push(e.data);
			};

			mediaRecorder.onstop = (e) => {
				console.log("Stop");
				assistantState = AssistantState.Processing;

				const blob = new Blob(chunks, {
					type: "audio/mpeg-3",
				});

				chunks = [];
				downloadHref = URL.createObjectURL(blob);
				assistantState = AssistantState.ReadyForInput;
			};
		});
	}

	function startStop() {
		if (!mediaRecorder) {
			return;
		}

		if (assistantState === AssistantState.Listening) {
			mediaRecorder.stop();
		} else if (assistantState === AssistantState.ReadyForInput) {
			mediaRecorder.start();
		}
	}
</script>

<svelte:head>
	<title>Cablulator</title>
</svelte:head>

<div class="wrapper">
	<h1>Calculate your cable's type</h1>

	<form>
		<label>
			<input
				type="radio"
				on:change={() =>
					(ampacityOrMaxPower = AmpacityOrMaxPower.Ampacity)}
				value={AmpacityOrMaxPower.Ampacity}
				checked={ampacityOrMaxPower === AmpacityOrMaxPower.Ampacity}
			/>
			Ampacity [A]
		</label>
		<input
			bind:value={ampacity}
			type="number"
			disabled={ampacityOrMaxPower !== AmpacityOrMaxPower.Ampacity}
		/>

		<label>
			<input
				type="radio"
				on:change={() =>
					(ampacityOrMaxPower = AmpacityOrMaxPower.MaxPower)}
				value={AmpacityOrMaxPower.MaxPower}
				checked={ampacityOrMaxPower === AmpacityOrMaxPower.MaxPower}
			/>
			Maximum active power of the load
		</label>
		<input
			type="number"
			bind:value={maxPower}
			disabled={ampacityOrMaxPower !== AmpacityOrMaxPower.MaxPower}
		/>
		<div class="gap" />

		<label for="veinsUnderLoad"> Number of veins under load </label>
		<input id="veinsUnderLoad" type="number" bind:value={veinsUnderLoad} />
		<div class="gap" />

		<label for="placement"> Placement </label>
		<select id="placement" bind:value={placements}>
			<option value={Placements.A1}>
				A1 (in a thermally insulated wall)
			</option>
			<option value={Placements.A2}>
				A2 (in pipes in a wall with very good insulation)
			</option>
			<option value={Placements.B2}>
				B2 (in pipes in a brick or concrete wall)
			</option>
			<option value={Placements.C}>
				C (under the plaster, on the wall, in full channels)
			</option>
			<option value={Placements.E}>
				E (in the air, in perforated channels)
			</option>
		</select>
		<div class="gap" />

		<label for="temperature"> Temperature of environment [°C] </label>
		<input id="temperature" type="number" bind:value={temperature} />
	</form>

	<div class="buttons-container">
		<button>
			<img src={calculate} alt="calculate" />
			<span> Calculate </span>
		</button>
		<button
			on:click={startStop}
			style="background-color: {assistantBtnColor};"
		>
			<img src={mic} alt="mic" />
			<span>
				{#if assistantState === AssistantState.ReadyForInput}
					Click to use an assistant
				{:else if assistantState === AssistantState.Listening}
					Listenning... Click to submit
				{:else}
					Processing...
				{/if}
			</span>
		</button>
	</div>
</div>

<style>
	.wrapper {
		display: flex;
		flex-direction: column;
		align-items: center;
		width: 100%;
	}

	h1 {
		font-family: "Roboto", sans-serif;
		color: rgb(255, 92, 0);
		font-size: 3rem;
		text-align: center;
	}

	button {
		width: 15rem;
		height: 3rem;
		margin: 1.5rem 1rem;
		background-color: rgb(0, 33, 95);
		display: flex;
		align-items: center;
		color: white;
		border: none;
		border-radius: 1.5rem;
		padding: 0 1rem;
	}

	button span {
		width: 100%;
		text-align: center;
	}

	.buttons-container {
		display: flex;
		flex-wrap: wrap;
		justify-content: center;
	}

	form {
		width: 100%;
		padding: 0 10vw;
	}

	label {
		display: block;
		width: 100%;
		font-size: 1.2rem;
		margin: 0.5rem;
	}

	input[type="number"],
	select {
		width: 100%;
		height: 2.5rem;
		border-radius: 2rem;
		border: rgb(0, 33, 95) solid;
		padding: 0 1rem;
		font-size: large;
	}

	input:disabled {
		background-color: lightgray;
	}

	.gap {
		width: 100%;
		height: 1rem;
	}
</style>
