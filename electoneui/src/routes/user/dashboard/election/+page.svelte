<script>
    import { onMount } from "svelte";
    import Navbar from "../../../../components/Navbar.svelte";
    import { goto } from "$app/navigation";

    let elections = [];
    let loading = true;
    let error = null;

    function hasElectionEnded(endDate) {
        return new Date(endDate) < new Date();
    }
    onMount(async () => {
        const token = localStorage.getItem("user_token");
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
                    class="card border-2 border-black"
                    on:click={() =>
                        !hasElectionEnded(e.end_date) &&
                        goto(`/user/dashboard/election/${e.election_id}`)}
                >
                    <div class="card-body">
                        <h2 class="card-title">{e.name}</h2>
                        <p class="text-sm text-gray-500">{e.location}</p>
                        <p class="text-sm text-gray-600">{e.description}</p>
                        <div class="mt-2 flex flex-wrap gap-2 items-center">
                            <span class="badge badge-outline"
                                >Start: {e.start_date}</span
                            >
                            <span class="badge badge-outline"
                                >End: {e.end_date}</span
                            >

                            {#if hasElectionEnded(e.end_date)}
                                <a
                                    href={`/results/${e.election_id}/`}
                                    class="badge bg-blue-500 text-white hover:bg-blue-600"
                                    on:click|stopPropagation
                                >
                                    View Election Analytics
                                </a>
                            {/if}
                        </div>
                    </div>
                </div>
            {/each}
        </div>
    {/if}
</section>
