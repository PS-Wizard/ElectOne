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
<section class="p-6 mx-auto space-y-8 max-w-2xl">
    <div
        class="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80"
        aria-hidden="true"
    >
        <div
            class="relative left-[calc(50%-11rem)] aspect-1155/678 w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-linear-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"
            style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
        ></div>
    </div>
    {#if error}
        <div class="alert alert-error"><span>{error}</span></div>
    {:else if !user || !citizenship || !voterCard}
        <div class="text-center text-lg text-gray-600">Loading...</div>
    {:else}
        <!-- Profile Picture -->
        <div class="flex justify-center">
            <img
                src={`http://localhost:3000${user.photos.split(",")[3]}`}
                alt="Profile Pic"
                class="rounded-full w-32 h-32 object-cover border-4 border-white shadow"
            />
        </div>

        <!-- Citizenship Accordion -->
        <details class="collapse collapse-arrow bg-base-200">
            <summary class="collapse-title text-xl font-medium"
                >Citizenship Details</summary
            >
            <div class="collapse-content space-y-2">
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
        </details>

        <!-- Voter Card Accordion -->
        <details class="collapse collapse-arrow bg-base-200">
            <summary class="collapse-title text-xl font-medium"
                >Voter Card</summary
            >
            <div class="collapse-content space-y-2">
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
        </details>

        <!-- Other Photos (Flex Column) -->
        <div class="space-y-4">
            <h2 class="text-xl font-bold">Documents</h2>
            <div class="flex flex-col gap-4">
                {#each user.photos
                    .split(",")
                    .filter((_, i) => i !== 3) as photo}
                    <img
                        src={`http://localhost:3000${photo}`}
                        alt="Document Photo"
                        class="rounded-xl shadow object-cover w-full max-h-96"
                    />
                {/each}
            </div>
        </div>
    {/if}
    <div
        class="absolute inset-x-0 top-[calc(100%-13rem)] -z-10 transform-gpu overflow-hidden blur-3xl sm:top-[calc(100%-30rem)]"
        aria-hidden="true"
    >
        <div
            class="relative left-[calc(50%+3rem)] aspect-1155/678 w-[36.125rem] -translate-x-1/2 bg-linear-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%+36rem)] sm:w-[72.1875rem]"
            style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
        ></div>
    </div>
</section>
