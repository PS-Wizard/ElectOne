<script>
    import { onMount } from "svelte";
    import Navbar from "../../../../components/Navbar.svelte";
    import { goto } from "$app/navigation";

    let elections = [];
    let loading = true;
    let error = null;

    onMount(async () => {
        const token = localStorage.getItem("token");
        if (!token) {
            error = "You're not logged in.";
            loading = false;
            return;
        }

        try {
            const res = await fetch(
                "http://localhost:3000/election?limit=10&offset=0",
                {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to fetch elections");
            const data = await res.json();
            elections = data ?? []; // make sure it's at least an empty array
        } catch (err) {
            error = err.message;
        } finally {
            loading = false;
        }
    });
</script>

<Navbar />
<section class="m-12">
    <h1 class="text-2xl font-bold mb-4">Elections</h1>

    {#if loading}
        <p class="text-gray-500">Loading elections...</p>
    {:else if error}
        <p class="text-red-500">Error: {error}</p>
    {:else if elections.length === 0}
        <p class="text-gray-500">There are no elections</p>
    {:else}
        <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-6 cursor-pointer">
            {#each elections as e}
                <div
                    class="card bg-base-100 shadow-xl"
                    on:click={() =>
                        goto(`/user/dashboard/election/${e.election_id}`)}
                >
                    <div class="card-body">
                        <h2 class="card-title">{e.name}</h2>
                        <p class="text-sm text-gray-500">{e.location}</p>
                        <p class="text-sm text-gray-600">{e.description}</p>
                        <div class="mt-2">
                            <span class="badge badge-outline"
                                >Start: {e.start_date}</span
                            >
                            <span class="badge badge-outline ml-2"
                                >End: {e.end_date}</span
                            >
                        </div>
                    </div>
                </div>
            {/each}
        </div>
    {/if}
</section>
