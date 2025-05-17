<script>
    import { onMount } from "svelte";
    import AdminNavbar from "../../../components/AdminNavbar.svelte";
    import Footer from "../../../components/footer.svelte";

    let appeals = [];
    let loading = true;
    let error = "";
    let search = "";
    let offset = 0;
    const limit = 10;
    let hasMore = true;

    let selectedPhoto = null;

    async function fetchAppeals() {
        loading = true;
        try {
            const res = await fetch(
                `http://localhost:3000/appeal?limit=${limit}&offset=${offset}`,
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to fetch appeals");
            const data = await res.json();
            appeals = data;
            hasMore = data.length === limit; // If we got less than limit, there are no more records
        } catch (err) {
            error = err.message;
        }
        loading = false;
    }

    async function approveAppeal(id) {
        try {
            const res = await fetch(
                `http://localhost:3000/appeal/${id}/approve`,
                {
                    method: "POST",
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to approve appeal");
            await fetchAppeals();
        } catch (err) {
            alert(err.message);
        }
    }

    async function rejectAppeal(id) {
        try {
            const res = await fetch(
                `http://localhost:3000/appeal/${id}/reject`,
                {
                    method: "POST",
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to reject appeal");
            await fetchAppeals();
        } catch (err) {
            alert(err.message);
        }
    }

    function nextPage() {
        offset += limit;
        fetchAppeals();
    }

    function prevPage() {
        if (offset >= limit) {
            offset -= limit;
        } else {
            offset = 0;
        }
        fetchAppeals();
    }

    onMount(fetchAppeals);
</script>

<AdminNavbar />
<section class="flex p-12 flex-col justify-center">
    <h1 class="text-center text-6xl uppercase font-medium tracking-wide">
        Handle Appeals
    </h1>

    <section class="flex justify-between items-center m-4">
        <div class="form-control w-full max-w-md">
            <input
                type="text"
                class="input rounded-lg w-full"
                bind:value={search}
                placeholder="Search appeals…"
            />
        </div>
    </section>

    <section class="flex justify-center m-4">
        {#if loading}
            <p>Loading appeals...</p>
        {:else if error}
            <p class="text-red-500">{error}</p>
        {:else}
            <div class="flex flex-col w-full">
                <div
                    class="overflow-x-auto rounded-box border border-base-content/5 bg-base-100 w-full"
                >
                    <table class="table w-full">
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Citizenship ID</th>
                                <th>Voter Card ID</th>
                                <th>Photos</th>
                                <th>Status</th>
                                <th>Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            {#each appeals as appeal}
                                {#if appeal.citizenship_id
                                    .toLowerCase()
                                    .includes(search.toLowerCase()) || appeal.voter_card_id
                                        .toLowerCase()
                                        .includes(search.toLowerCase())}
                                    <tr>
                                        <td>{appeal.appeal_id}</td>
                                        <td>{appeal.citizenship_id}</td>
                                        <td>{appeal.voter_card_id}</td>
                                        <td>
                                            <div class="flex gap-1">
                                                {#each appeal.photos.split(",") as photo}
                                                    <img
                                                        src={photo}
                                                        alt="Appeal Photo"
                                                        class="w-12 h-12 object-cover rounded-full cursor-pointer hover:scale-105 transition"
                                                        on:click={() =>
                                                            (selectedPhoto = photo)}
                                                    />
                                                {/each}
                                            </div>
                                        </td>
                                        <td>{appeal.status}</td>
                                        <td
                                            class="p-4 flex gap-2 justify-center"
                                        >
                                            <button
                                                class="btn btn-sm btn-primary"
                                                disabled={appeal.status !==
                                                    "Pending"}
                                                on:click={() =>
                                                    approveAppeal(
                                                        appeal.appeal_id,
                                                    )}
                                            >
                                                Approve
                                            </button>
                                            <button
                                                class="btn btn-sm btn-ghost"
                                                disabled={appeal.status !==
                                                    "Pending"}
                                                on:click={() =>
                                                    rejectAppeal(
                                                        appeal.appeal_id,
                                                    )}
                                            >
                                                Reject
                                            </button>
                                        </td>
                                    </tr>
                                {/if}
                            {/each}
                        </tbody>
                    </table>
                </div>

                <!-- Pagination controls -->
                <div class="flex justify-between items-center mt-4">
                    <button
                        class="btn btn-sm"
                        disabled={offset === 0}
                        on:click={prevPage}
                    >
                        Previous
                    </button>
                    <button
                        class="btn btn-sm"
                        disabled={!hasMore}
                        on:click={nextPage}
                    >
                        Next
                    </button>
                </div>
            </div>
        {/if}
    </section>
</section>

{#if selectedPhoto}
    <div
        class="fixed inset-0 bg-black/80 flex items-center justify-center z-50"
    >
        <div class="relative">
            <img
                src={selectedPhoto}
                class="max-w-full max-h-screen rounded-lg shadow-lg"
            />
            <button
                class="absolute top-2 right-2 btn btn-sm btn-circle btn-error"
                on:click={() => (selectedPhoto = null)}
            >
                ✕
            </button>
        </div>
    </div>
{/if}

<style>
    /* Optional: Add some custom styles for very small screens if needed */
    @media (max-width: 640px) {
        .filter-section {
            flex-direction: column;
            align-items: stretch;
        }

        .filter-section > div {
            width: 100%;
        }

        .flex-col.sm\:flex-row > div {
            width: 100%;
        }

        .flex-col.sm\:flex-row {
            align-items: stretch;
        }

        .table thead th {
            padding: 0.75rem 0.5rem;
            font-size: 0.8rem;
        }

        .table tbody td {
            padding: 0.75rem 0.5rem;
            font-size: 0.9rem;
        }

        .text-right.font-semibold.uppercase.tracking-wider.text-gray-600.py-3.px-4 {
            text-align: center;
        }

        .table tbody tr td:last-child {
            text-align: center;
        }

        .space-x-2.flex.justify-end {
            justify-content: center !important;
        }
    }
</style>

