<script lang="ts">
    import { onMount } from "svelte";

    let appeals = [];

    onMount(async () => {
        await fetchAppeals();
    });

    async function fetchAppeals() {
        try {
            const res = await fetch(
                "https://localhost:3000/api/secure/appeals/",
            );
            if (!res.ok) throw new Error("Failed to fetch appeals");
            appeals = await res.json();
        } catch (err) {
            console.error("Error fetching appeals:", err);
        }
    }

    async function approveAppeal(id: number) {
        try {
            const res = await fetch(
                `https://localhost:3000/api/secure/appeals/${id}`,
                {
                    method: "POST",
                },
            );
            if (!res.ok) throw new Error("Failed to approve appeal");
            await fetchAppeals(); // refresh after approve
        } catch (err) {
            console.error("Error approving appeal:", err);
        }
    }

    async function rejectAppeal(id: number) {
        try {
            const res = await fetch(
                `https://localhost:3000/api/secure/appeals/${id}`,
                {
                    method: "DELETE",
                },
            );
            if (!res.ok) throw new Error("Failed to reject appeal");
            await fetchAppeals(); // refresh after reject
        } catch (err) {
            console.error("Error rejecting appeal:", err);
        }
    }
</script>

<main class="p-4">
    <h1 class="text-2xl font-bold mb-4">Pending Appeals</h1>

    {#if appeals.length > 0}
        {#each appeals as appeal}
            <div class="border p-4 rounded mb-4 shadow">
                <h2 class="text-xl font-semibold">{appeal.fullName}</h2>
                <p><strong>Citizen ID:</strong> {appeal.citizenID}</p>
                <p><strong>DOB:</strong> {appeal.dateOfBirth}</p>
                <p><strong>Residence:</strong> {appeal.placeOfResidence}</p>
                <p><strong>Phone:</strong> {appeal.phoneNumber}</p>
                <p><strong>Tags:</strong> {appeal.tags}</p>
                <div>
                    <strong>Photos:</strong>
                    <ul class="list-disc ml-6">
                        {#each appeal.photos as photo}
                            <li>{photo}</li>
                        {/each}
                    </ul>
                </div>

                <div class="mt-4 flex gap-2">
                    <button
                        on:click={() => approveAppeal(appeal.appealID)}
                        class="bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded"
                    >
                        Approve
                    </button>
                    <button
                        on:click={() => rejectAppeal(appeal.appealID)}
                        class="bg-red-500 hover:bg-red-600 text-white px-4 py-2 rounded"
                    >
                        Reject
                    </button>
                </div>
            </div>
        {/each}
    {:else}
        <p>No pending appeals.</p>
    {/if}
</main>

<style>
    main {
        max-width: 800px;
        margin: auto;
    }
</style>
