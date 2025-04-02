<script>
    import { onMount } from "svelte";
    import { GetUsers, DeleteUser, UpdateUser, CreateUser } from "./api.js";

    let users = [];
    let searchQuery = "";
    let newUser = { citizenID: "", password: "", tags: "", phoneNumber: "" };
    let editingUser = null;

    async function loadUsers() {
        users = await GetUsers();
    }

    async function removeUser(id) {
        await DeleteUser(id);
        loadUsers();
    }

    async function addUser() {
        await CreateUser(newUser);
        newUser = { citizenID: "", password: "", tags: "", phoneNumber: "" };
        loadUsers();
    }

    async function saveUser() {
        if (editingUser) {
            await UpdateUser(editingUser.userID, editingUser);
            editingUser = null;
            loadUsers();
        }
    }

    onMount(loadUsers);
</script>

<main class="mx-auto p-10 bg-black text-gray-300">
    <div class="flex justify-between items-center mb-4">
        <input
            type="text"
            bind:value={searchQuery}
            placeholder="Search by Citizen ID..."
            class="input input-bordered bg-[#0a0a0a] text-white"
        />
        <button class="btn btn-primary" on:click={() => showModal.showModal()}
            >+ New Record</button
        >
    </div>

    <div class="overflow-x-auto flex justify-center">
        <table
            class="table w-full border border-[#070707] bg-[#0a0a0a] text-white"
        >
            <thead>
                <tr class="bg-[#070707]">
                    <th class="text-center">User ID</th>
                    <th class="text-center">Citizen ID</th>
                    <th class="text-center">Phone Number</th>
                    <th class="text-center">Tags</th>
                    <th class="text-center">Password</th>
                    <th class="text-center">Actions</th>
                </tr>
            </thead>
            <tbody>
                {#each users.filter((u) => u.citizenID
                        .toLowerCase()
                        .includes(searchQuery.toLowerCase())) as user}
                    <tr class="hover">
                        {#if editingUser?.userID === user.userID}
                            <td
                                ><input
                                    bind:value={editingUser.userID}
                                    disabled
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td
                                ><input
                                    bind:value={editingUser.citizenID}
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td
                                ><input
                                    bind:value={editingUser.phoneNumber}
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td
                                ><input
                                    bind:value={editingUser.tags}
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td
                                ><input
                                    bind:value={editingUser.password}
                                    type="password"
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td>
                                <button
                                    class="btn btn-primary"
                                    on:click={saveUser}>Save</button
                                >
                            </td>
                        {:else}
                            <td class="text-center border-r border-white"
                                >{user.userID}</td
                            >
                            <td class="text-center border-r border-white"
                                >{user.citizenID}</td
                            >
                            <td class="text-center border-r border-white"
                                >{user.phoneNumber}</td
                            >
                            <td class="text-center border-r border-white"
                                >{user.tags}</td
                            >
                            <td class="text-center border-r border-white"
                                >********</td
                            >
                            <!-- Masked password -->
                            <td>
                                <button
                                    class="btn btn-secondary text-black"
                                    on:click={() => (editingUser = { ...user })}
                                    >Edit</button
                                >
                                <button
                                    class="btn btn-error"
                                    on:click={() => removeUser(user.userID)}
                                    >Delete</button
                                >
                            </td>
                        {/if}
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>

    <dialog id="showModal" class="modal">
        <div class="modal-box bg-[#0a0a0a] text-white">
            <h2 class="text-xl mb-4">Add New User</h2>
            <input
                bind:value={newUser.citizenID}
                placeholder="Citizen ID"
                class="input input-bordered w-full bg-black text-white mb-2"
            />
            <input
                bind:value={newUser.password}
                placeholder="Password"
                type="password"
                class="input input-bordered w-full bg-black text-white mb-2"
            />
            <input
                bind:value={newUser.phoneNumber}
                placeholder="Phone Number"
                class="input input-bordered w-full bg-black text-white mb-2"
            />
            <input
                bind:value={newUser.tags}
                placeholder="Tags"
                class="input input-bordered w-full bg-black text-white mb-4"
            />

            <div class="flex justify-end space-x-2">
                <form method="dialog">
                    <button class="btn btn-secondary">Cancel</button>
                </form>
                <button class="btn btn-primary" on:click={addUser}
                    >Create</button
                >
            </div>
        </div>
    </dialog>
</main>
