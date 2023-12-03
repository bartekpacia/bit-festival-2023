<script lang="ts">
    import type { GoRsp } from "$lib/scripts/goBackend";

    export let result: GoRsp;
</script>

<div class="wrapper">
    <div class="row">
        <h1>Results:</h1>
    </div>
    <div class="row">
        <span>cableType:</span>
        <span>{result.cableType}</span>
    </div>
    <div class="row">
        <span>veinsNumber:</span>
        <span>{result.veinsNumber}</span>
    </div>
    <div class="row">
        <span>crossSection:</span>
        <span>{result.crossSection}</span>
    </div>
    <div class="row">
        <span>insulationType:</span>
        <span>{result.insulationType}</span>
    </div>
    <div class="row">
        <span>placements:</span>
        <span>{result.placements}</span>
    </div>
    <div class="row">
        <span>temperature:</span>
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
