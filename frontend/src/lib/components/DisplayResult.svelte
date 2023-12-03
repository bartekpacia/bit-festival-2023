<script lang="ts">
    import type { GoRsp } from "$lib/scripts/goBackend";

    export let result: GoRsp;
</script>

<div class="wrapper">
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
</div>
