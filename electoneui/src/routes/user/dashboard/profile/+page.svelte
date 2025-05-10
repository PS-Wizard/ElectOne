<script>
    import { onMount } from "svelte";
    import Navbar from "../../../../components/Navbar.svelte";

    let user = null;
    let citizenship = null;
    let voterCard = null;
    let error = null;

    onMount(async () => {
        try {
            const token = localStorage.getItem("user_token");

            const userRes = await fetch("http://localhost:3000/user/0", {
                headers: {
                    Authorization: `Bearer ${token}`,
                    Accept: "application/json",
                },
            });
            if (!userRes.ok) throw new Error("Failed to fetch user");
            user = await userRes.json();

            const citizenRes = await fetch(
                `http://localhost:3000/citizen/${user.citizenship_id}`,
                {
                    headers: {
                        Authorization: `Bearer ${token}`,
                        Accept: "application/json",
                    },
                },
            );
            if (!citizenRes.ok) throw new Error("Failed to fetch citizenship");
            citizenship = await citizenRes.json();

            const voterCardRes = await fetch(
                `http://localhost:3000/voter_card/${user.voter_card_id}`,
                {
                    headers: {
                        Authorization: `Bearer ${token}`,
                        Accept: "application/json",
                    },
                },
            );
            if (!voterCardRes.ok) throw new Error("Failed to fetch voter card");
            voterCard = await voterCardRes.json();
        } catch (err) {
            error = err.message;
        }
    });
</script>

<Navbar />
<section class="p-6 mx-auto space-y-8">
    {#if error}
        <div class="alert alert-error">
            <span>{error}</span>
        </div>
    {:else if !user || !citizenship || !voterCard}
        <div class="text-center text-lg text-gray-600">Loading...</div>
    {:else}
        <!-- User Info -->
        <div class="card p-6">
            <div class="flex justify-center gap-2">
                {#each user.photos.split(",") as photo}
                    <img
                        src={`http://localhost:3000${photo}`}
                        alt="User Photo"
                        class="rounded-xl shadow-sm object-cover aspect-[4/5]"
                    />
                {/each}
            </div>
        </div>

        <div class="card p-6">
            <h2 class="text-2xl font-bold uppercase tracking-wide mb-4">Citizenship Details</h2>
            <div class="grid sm:grid-cols-2 gap-4">
                <p>
                    <span class="font-medium">Full Name:</span>
                    {citizenship.full_name}
                </p>
                <p>
                    <span class="font-medium">Date of Birth:</span>
                    {citizenship.date_of_birth}
                </p>
                <p>
                    <span class="font-medium">Birth Place:</span>
                    {citizenship.birth_place}
                </p>
                <p>
                    <span class="font-medium">Permanent Address:</span>
                    {citizenship.permanent_address}
                </p>
            </div>
        </div>

        <!-- Voter Card Info -->
        <div class="card p-6">
            <h2 class="text-2xl font-bold mb-4">Voter Card</h2>

            <div class="grid sm:grid-cols-2 gap-4">
                <p>
                    <span class="font-medium">Voter Card ID:</span>
                    {voterCard.voter_card_id}
                </p>
                <p>
                    <span class="font-medium">Linked Citizenship:</span>
                    {voterCard.citizenship_id}
                </p>
                <p>
                    <span class="font-medium">Location:</span>
                    {voterCard.location}
                </p>
            </div>
        </div>
    {/if}
</section>
