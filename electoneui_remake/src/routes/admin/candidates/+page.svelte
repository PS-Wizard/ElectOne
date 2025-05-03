<script>
    import { onMount } from "svelte";
    let candidates = [];
    let loading = true;
    let error = "";

    let newCandidate = {
        citizen_id: "",
        election_id: "",
        profile_path: "",
        bio: "",
        post: "",
    };

    let editingCandidate = null;

    let newCandidateModal;
    let editCandidateModal;

    async function fetchCandidates() {
        loading = true;
        try {
            const res = await fetch(
                "http://localhost:3000/candidate?limit=100&offset=0",
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("token")}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to fetch candidates");
            candidates = await res.json();
        } catch (err) {
            error = err.message;
        }
        loading = false;
    }

    async function createCandidate() {
        try {
            const payload = {
                ...newCandidate,
                election_id: Number(newCandidate.election_id),
            };

            const res = await fetch("http://localhost:3000/candidate", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${localStorage.getItem("token")}`,
                },
                body: JSON.stringify(payload),
            });
            if (!res.ok) throw new Error("Failed to create candidate");
            await fetchCandidates();
            newCandidate = {
                citizen_id: "",
                election_id: "",
                profile_path: "",
                bio: "",
                post: "",
            };
            newCandidateModal.close();
        } catch (err) {
            alert(err.message);
        }
    }

    function openEditModal(candidate) {
        editingCandidate = { ...candidate };
        editCandidateModal.showModal();
    }

    async function updateCandidate() {
        if (!editingCandidate) return;
        try {
            const payload = {
                ...editingCandidate,
                election_id: Number(editingCandidate.election_id),
            };

            const res = await fetch(
                `http://localhost:3000/candidate/${editingCandidate.candidate_id}`,
                {
                    method: "PUT",
                    headers: {
                        "Content-Type": "application/json",
                        Authorization: `Bearer ${localStorage.getItem("token")}`,
                    },
                    body: JSON.stringify(payload),
                },
            );
            if (!res.ok) throw new Error("Failed to update candidate");
            await fetchCandidates();
            editingCandidate = null;
            editCandidateModal.close();
        } catch (err) {
            alert(err.message);
        }
    }

    async function deleteCandidate(id) {
        if (!confirm("Delete this candidate?")) return;
        try {
            const res = await fetch(`http://localhost:3000/candidate/${id}`, {
                method: "DELETE",
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("token")}`,
                },
            });
            if (!res.ok) throw new Error("Failed to delete");
            await fetchCandidates();
        } catch (err) {
            alert(err.message);
        }
    }

    onMount(fetchCandidates);
</script>

<button
    class="btn btn-primary mb-4"
    on:click={() => newCandidateModal.showModal()}>New Candidate</button
>

<dialog bind:this={newCandidateModal} class="modal">
    <div class="modal-box">
        <h3 class="font-bold text-lg">New Candidate</h3>
        <div class="py-2 flex flex-col gap-2">
            <input
                class="input input-bordered"
                placeholder="Citizen ID"
                bind:value={newCandidate.citizen_id}
            />
            <input
                class="input input-bordered"
                placeholder="Election ID"
                bind:value={newCandidate.election_id}
            />
            <input
                class="input input-bordered"
                placeholder="Profile Path"
                bind:value={newCandidate.profile_path}
            />
            <input
                class="input input-bordered"
                placeholder="Bio"
                bind:value={newCandidate.bio}
            />
            <input
                class="input input-bordered"
                placeholder="Post"
                bind:value={newCandidate.post}
            />
        </div>
        <div class="modal-action">
            <form method="dialog" class="flex gap-2">
                <button
                    type="button"
                    class="btn btn-success"
                    on:click={createCandidate}>Create</button
                >
                <button class="btn">Cancel</button>
            </form>
        </div>
    </div>
</dialog>

<dialog bind:this={editCandidateModal} class="modal">
    <div class="modal-box">
        <h3 class="font-bold text-lg">Edit Candidate</h3>
        {#if editingCandidate}
            <div class="py-2 flex flex-col gap-2">
                <input
                    class="input input-bordered"
                    value={editingCandidate.candidate_id}
                    readonly
                />
                <input
                    class="input input-bordered"
                    placeholder="Citizen ID"
                    bind:value={editingCandidate.citizen_id}
                />
                <input
                    class="input input-bordered"
                    placeholder="Election ID"
                    bind:value={editingCandidate.election_id}
                />
                <input
                    class="input input-bordered"
                    placeholder="Profile Path"
                    bind:value={editingCandidate.profile_path}
                />
                <input
                    class="input input-bordered"
                    placeholder="Bio"
                    bind:value={editingCandidate.bio}
                />
                <input
                    class="input input-bordered"
                    placeholder="Post"
                    bind:value={editingCandidate.post}
                />
            </div>
            <div class="modal-action">
                <form method="dialog" class="flex gap-2">
                    <button
                        type="button"
                        class="btn btn-warning"
                        on:click={updateCandidate}>Update</button
                    >
                    <button class="btn">Cancel</button>
                </form>
            </div>
        {/if}
    </div>
</dialog>

{#if loading}
    <p>Loading...</p>
{:else if error}
    <p class="text-red-500">{error}</p>
{:else}
    <div class="overflow-x-auto">
        <table class="table w-full">
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Citizen ID</th>
                    <th>Election ID</th>
                    <th>Profile Path</th>
                    <th>Bio</th>
                    <th>Post</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                {#each candidates as c}
                    <tr>
                        <td>{c.candidate_id}</td>
                        <td>{c.citizen_id}</td>
                        <td>{c.election_id}</td>
                        <td>{c.profile_path}</td>
                        <td>{c.bio}</td>
                        <td>{c.post}</td>
                        <td class="flex gap-2">
                            <button
                                class="btn btn-sm btn-warning"
                                on:click={() => openEditModal(c)}>Edit</button
                            >
                            <button
                                class="btn btn-sm btn-error"
                                on:click={() => deleteCandidate(c.candidate_id)}
                                >Delete</button
                            >
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>
{/if}

<style>
    .modal-box {
        max-width: 560px;
        width: 90%;
    }
</style>
