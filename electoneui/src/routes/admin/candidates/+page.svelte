<script>
    import { onMount } from "svelte";
    import AdminNavbar from "../../../components/AdminNavbar.svelte";

    let candidates = [];
    let loading = true;
    let error = "";
    let search = "";
    let selectedPhoto = "";
    let newCandidate = {
        citizen_id: "",
        election_id: "",
        bio: "",
        post: "",
        photo: null, // Added for file upload
    };

    let editingCandidate = null;
    let editedPhoto = null; // To hold the updated photo file

    let newCandidateModal;
    let editCandidateModal;

    async function fetchCandidates() {
        loading = true;
        try {
            const res = await fetch(
                "http://localhost:3000/candidate?limit=100&offset=0",
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
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
        const formData = new FormData();
        formData.append("citizenship_id", newCandidate.citizen_id);
        formData.append("election_id", newCandidate.election_id);
        formData.append("bio", newCandidate.bio);
        formData.append("post", newCandidate.post);
        formData.append("photo", newCandidate.photo); // Adding the photo to the form data

        try {
            const res = await fetch("http://localhost:3000/candidate", {
                method: "POST",
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                },
                body: formData,
            });
            if (!res.ok) throw new Error("Failed to create candidate");
            await fetchCandidates();
            newCandidate = {
                citizen_id: "",
                election_id: "",
                bio: "",
                post: "",
                photo: null, // Reset photo after submission
            };
            newCandidateModal.close();
        } catch (err) {
            alert(err.message);
        }
    }

    function openEditModal(candidate) {
        editingCandidate = { ...candidate };
        editedPhoto = null; // Reset edited photo when opening modal
        editCandidateModal.showModal();
    }

    async function updateCandidate() {
        if (!editingCandidate) return;
        const formData = new FormData();
        formData.append("citizenship_id", editingCandidate.citizen_id);
        formData.append("election_id", editingCandidate.election_id);
        formData.append("bio", editingCandidate.bio);
        formData.append("post", editingCandidate.post);
        if (editedPhoto) {
            formData.append("photo", editedPhoto); // Include new photo if selected
        }

        try {
            const res = await fetch(
                `http://localhost:3000/candidate/${editingCandidate.candidate_id}`,
                {
                    method: "PUT",
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                    },
                    body: formData,
                },
            );
            if (!res.ok) throw new Error("Failed to update candidate");
            await fetchCandidates();
            editingCandidate = null;
            editedPhoto = null;
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
                    Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                },
            });
            if (!res.ok) throw new Error("Failed to delete candidate");
            await fetchCandidates();
        } catch (err) {
            alert(err.message);
        }
    }

    onMount(fetchCandidates);
</script>

<AdminNavbar />
<section class="flex p-12 flex-col justify-center">
    <h1 class="text-center text-6xl uppercase font-medium tracking-wide">
        Manage Candidates
    </h1>

    <section class="flex justify-between items-center m-4">
        <div class="form-control w-full max-w-md">
            <input
                type="text"
                class="input rounded-lg w-full"
                placeholder="Search candidates…"
                bind:value={search}
            />
        </div>
        <button
            class="btn btn-primary ml-4 rounded-lg"
            on:click={() => newCandidateModal.showModal()}
        >
            + New Candidate
        </button>
    </section>

    <section class="flex justify-center m-4">
        {#if loading}
            <p>Loading candidates...</p>
        {:else if error}
            <p class="text-red-500">{error}</p>
        {:else}
            <div
                class="overflow-x-auto rounded-box border border-base-content/5 bg-base-100 w-full"
            >
                <table class="table w-full">
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>Citizen ID</th>
                            <th>Election ID</th>
                            <th>Profile</th>
                            <th>Bio</th>
                            <th>Post</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each candidates as c}
                            {#if c.citizen_id
                                .toLowerCase()
                                .includes(search.toLowerCase()) || c.election_id
                                    .toString()
                                    .includes(search)}
                                <tr>
                                    <td>{c.candidate_id}</td>
                                    <td>{c.citizen_id}</td>
                                    <td>{c.election_id}</td>
                                    <td>
                                        <img
                                            src={`http://localhost:3000${c.candidate_photo}`}
                                            alt="Candidate Photo"
                                            class="w-12 h-12 object-cover rounded-full cursor-pointer hover:scale-105 transition"
                                            on:click={() =>
                                                (selectedPhoto = `http://localhost:3000${c.candidate_photo}`)}
                                        />
                                    </td>
                                    <td>{c.bio}</td>
                                    <td>{c.post}</td>
                                    <td class="flex gap-2">
                                        <button
                                            class="btn btn-sm btn-warning"
                                            on:click={() => openEditModal(c)}
                                        >
                                            Edit
                                        </button>
                                        <button
                                            class="btn btn-sm btn-error"
                                            on:click={() =>
                                                deleteCandidate(c.candidate_id)}
                                        >
                                            Delete
                                        </button>
                                    </td>
                                </tr>
                            {/if}
                        {/each}
                    </tbody>
                </table>
            </div>
        {/if}
    </section>
</section>

<dialog bind:this={newCandidateModal} class="modal">
    <div class="modal-box">
        <h3 class="font-bold text-lg">Create New Candidate</h3>
        <div class="py-2 flex flex-col gap-2">
            <input
                class="input input-bordered w-full rounded-lg"
                placeholder="Citizen ID"
                bind:value={newCandidate.citizen_id}
                required
            />
            <input
                class="input input-bordered w-full rounded-lg"
                placeholder="Election ID"
                bind:value={newCandidate.election_id}
                required
            />
            <label for="new_candidate_photo" class="label">
                <span class="label-text">Profile Photo</span>
            </label>
            <input
                type="file"
                id="new_candidate_photo"
                class="file-input file-input-bordered w-full rounded-lg"
                accept="image/*"
                on:change={(e) => (newCandidate.photo = e.target.files[0])}
                required
            />
            <textarea
                class="textarea textarea-bordered w-full rounded-lg"
                placeholder="Bio"
                bind:value={newCandidate.bio}
                required
            />
            <input
                class="input input-bordered w-full rounded-lg"
                placeholder="Post"
                bind:value={newCandidate.post}
                required
            />
        </div>
        <div class="modal-action">
            <form method="dialog" class="flex gap-2">
                <button
                    type="button"
                    class="btn btn-primary rounded-lg"
                    on:click={createCandidate}>Create</button
                >
                <button class="btn btn-outline rounded-lg">Cancel</button>
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
                <label for="edit_candidate_photo" class="label">
                    <span class="label-text"
                        >Update Profile Photo (Optional)</span
                    >
                </label>
                <input
                    type="file"
                    id="edit_candidate_photo"
                    class="file-input file-input-bordered w-full rounded-lg"
                    accept="image/*"
                    on:change={(e) => (editedPhoto = e.target.files[0])}
                />
                <textarea
                    class="textarea textarea-bordered w-full rounded-lg"
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
    .modal-box {
        max-width: 560px;
        width: 90%;
    }
</style>
