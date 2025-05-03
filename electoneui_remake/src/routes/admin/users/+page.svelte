<script>
    import { onMount } from "svelte";

    let users = [];
    let loading = true;
    let error = "";

    let newUser = {
        citizenship_id: "",
        voter_card_id: "",
        password: "",
        totp_secret: "",
        photos: "",
    };

    let editingUser = null;
    let newUserModal;
    let editUserModal;

    async function fetchUsers() {
        loading = true;
        try {
            const res = await fetch(
                "http://localhost:3000/user?limit=100&offset=0",
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("token")}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to fetch users");
            users = await res.json();
        } catch (err) {
            error = err.message;
        }
        loading = false;
    }

    async function createUser() {
        try {
            const res = await fetch("http://localhost:3000/user", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${localStorage.getItem("token")}`,
                },
                body: JSON.stringify(newUser),
            });
            if (!res.ok) throw new Error("Failed to create user");
            await fetchUsers();
            newUser = {
                citizenship_id: "",
                voter_card_id: "",
                password: "",
                totp_secret: "",
                photos: "",
            };
            newUserModal.close();
        } catch (err) {
            alert(err.message);
        }
    }

    function openEditModal(user) {
        editingUser = { ...user };
        editUserModal.showModal();
    }

    async function updateUser() {
        if (!editingUser) return;
        try {
            const res = await fetch(
                `http://localhost:3000/user/${editingUser.user_id}`,
                {
                    method: "PUT",
                    headers: {
                        "Content-Type": "application/json",
                        Authorization: `Bearer ${localStorage.getItem("token")}`,
                    },
                    body: JSON.stringify(editingUser),
                },
            );
            if (!res.ok) throw new Error("Failed to update user");
            await fetchUsers();
            editingUser = null;
            editUserModal.close();
        } catch (err) {
            alert(err.message);
        }
    }

    async function deleteUser(id) {
        if (!confirm("Delete this user?")) return;
        try {
            const res = await fetch(`http://localhost:3000/user/${id}`, {
                method: "DELETE",
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("token")}`,
                },
            });
            if (!res.ok) throw new Error("Failed to delete user");
            await fetchUsers();
        } catch (err) {
            alert(err.message);
        }
    }

    onMount(fetchUsers);
</script>

<button class="btn btn-primary mb-4" on:click={() => newUserModal.showModal()}
    >New User</button
>

<dialog bind:this={newUserModal} id="new_user_modal" class="modal">
    <div class="modal-box">
        <h3 class="font-bold text-lg">New User</h3>
        <div class="py-2 flex flex-col gap-2">
            <input
                class="input input-bordered"
                placeholder="Citizenship ID"
                bind:value={newUser.citizenship_id}
            />
            <input
                class="input input-bordered"
                placeholder="Voter Card ID"
                bind:value={newUser.voter_card_id}
            />
            <input
                type="password"
                class="input input-bordered"
                placeholder="Password"
                bind:value={newUser.password}
            />
            <input
                class="input input-bordered"
                placeholder="TOTP Secret"
                bind:value={newUser.totp_secret}
            />
            <input
                class="input input-bordered"
                placeholder="Photos (JSON)"
                bind:value={newUser.photos}
            />
        </div>
        <div class="modal-action">
            <form method="dialog" class="flex gap-2">
                <button
                    type="button"
                    class="btn btn-success"
                    on:click={createUser}>Create</button
                >
                <button class="btn">Cancel</button>
            </form>
        </div>
    </div>
</dialog>

<dialog bind:this={editUserModal} id="edit_user_modal" class="modal">
    <div class="modal-box">
        <h3 class="font-bold text-lg">Edit User</h3>
        {#if editingUser}
            <div class="py-2 flex flex-col gap-2">
                <label class="label">
                    <span class="label-text">Citizenship ID</span>
                </label>
                <input
                    type="text"
                    class="input input-bordered"
                    bind:value={editingUser.citizenship_id}
                />
                <label class="label">
                    <span class="label-text">Voter Card ID</span>
                </label>
                <input
                    type="text"
                    class="input input-bordered"
                    bind:value={editingUser.voter_card_id}
                />
                <label class="label">
                    <span class="label-text">Password</span>
                </label>
                <input
                    type="password"
                    class="input input-bordered"
                    placeholder="Password"
                    bind:value={editingUser.password}
                />
                <label class="label">
                    <span class="label-text">TOTP Secret</span>
                </label>
                <input
                    type="text"
                    class="input input-bordered"
                    bind:value={editingUser.totp_secret}
                />
                <label class="label">
                    <span class="label-text">Photos (JSON)</span>
                </label>
                <input
                    type="text"
                    class="input input-bordered"
                    bind:value={editingUser.photos}
                />
            </div>
            <div class="modal-action">
                <form method="dialog" class="flex gap-2">
                    <button
                        type="button"
                        class="btn btn-warning"
                        on:click={updateUser}>Update</button
                    >
                    <button class="btn">Cancel</button>
                </form>
            </div>
        {/if}
    </div>
</dialog>

{#if loading}
    <p>Loading users...</p>
{:else if error}
    <p class="text-red-500">{error}</p>
{:else}
    <div class="overflow-x-auto">
        <table class="table w-full">
            <thead>
                <tr>
                    <th>User ID</th>
                    <th>Citizenship ID</th>
                    <th>Voter Card ID</th>
                    <th>TOTP Secret</th>
                    <th>Photos</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                {#each users as user}
                    <tr>
                        <td>{user.user_id}</td>
                        <td>{user.citizenship_id}</td>
                        <td>{user.voter_card_id}</td>
                        <td>{user.totp_secret}</td>
                        <td>{user.photos}</td>
                        <td class="flex gap-2">
                            <button
                                class="btn btn-sm btn-warning"
                                on:click={() => openEditModal(user)}
                                >Edit</button
                            >
                            <button
                                class="btn btn-sm btn-error"
                                on:click={() => deleteUser(user.user_id)}
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
    /* Optional: Add some styling for better visual appearance */
    .modal-box {
        max-width: 560px; /* Adjust as needed */
        width: 90%;
    }
</style>
