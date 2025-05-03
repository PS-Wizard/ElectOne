<script>
    import { onMount } from "svelte";

    let appeals = [];
    let loading = true;
    let error = "";

    async function fetchAppeals() {
        loading = true;
        try {
            const res = await fetch(
                "http://localhost:3000/appeal?limit=100&offset=0",
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("token")}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to fetch appeals");
            appeals = await res.json();
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
                        Authorization: `Bearer ${localStorage.getItem("token")}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to approve appeal");
            await fetchAppeals(); // Refresh the list after approval
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
                        Authorization: `Bearer ${localStorage.getItem("token")}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to reject appeal");
            await fetchAppeals(); // Refresh the list after rejection
        } catch (err) {
            alert(err.message);
        }
    }

    onMount(fetchAppeals);
</script>

{#if loading}
    <p>Loading appeals...</p>
{:else if error}
    <p class="text-red-500">{error}</p>
{:else}
    <div class="overflow-x-auto">
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
                    <tr>
                        <td>{appeal.appeal_id}</td>
                        <td>{appeal.citizenship_id}</td>
                        <td>{appeal.voter_card_id}</td>
                        <td>{appeal.photos}</td>
                        <td>{appeal.status}</td>
                        <td class="flex gap-2">
                            <button
                                class="btn btn-sm btn-success"
                                disabled={appeal.status !== "Pending"}
                                on:click={() => approveAppeal(appeal.appeal_id)}
                                >Approve</button
                            >
                            <button
                                class="btn btn-sm btn-error"
                                disabled={appeal.status !== "Pending"}
                                on:click={() => rejectAppeal(appeal.appeal_id)}
                                >Reject</button
                            >
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>
{/if}
