<script lang="ts">
    import type { GoRsp } from "$lib/scripts/goBackend";

    export let result: GoRsp;
</script>

<div class="wrapper">
    <div class="row">
        <h1>Results:</h1>
    </div>
    <div class="row">
        <span>Cable type:</span>
        <span>{result.cableType}</span>
    </div>
    <div class="row">
        <span>Veins number:</span>
        <span>{result.veinsNumber}</span>
    </div>
    <div class="row">
        <span>Cross-section:</span>
        <span>{result.crossSection}</span>
    </div>
    <div class="row">
        <span>Insulation type:</span>
        <span>{result.insulationType}</span>
    </div>
    <div class="row">
        <span>Placements:</span>
        <span>{result.placements}</span>
    </div>
    <div class="row">
        <span>Temperature:</span>
        <span>{result.temperature}</span>
    </div>
    <div class="row link">
        <a
            href="https://www.nkt.com.pl/produkty-rozwiazania/niskie-napiecie/przewody-instalacyjne/ydy-450-750-v"
            target="_blank"
        >
            Click here to find this cable
        </a>
    </div>
</div>

<style>
    .wrapper {
        width: 100%;
        padding: 10vh 20vw;
    }

    .row {
        font-size: 1.2rem;
        display: flex;
        justify-content: space-between;
        margin-top: 0.5rem;
    }

    .link {
        margin-top: 3rem;
        justify-content: center;
    }
</style>
