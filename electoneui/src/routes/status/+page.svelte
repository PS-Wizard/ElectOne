<script>
    import Footer from "../../components/footer.svelte";
    import Navbar from "../../components/Navbar.svelte";

    let appeal_id = "";
    let citizenship_id = "";
    let voter_card_id = "";
    let password = "";

    let loading = false;
    let result = null;
    let error = null;

    async function checkStatus() {
        loading = true;
        result = null;
        error = null;

        try {
            const res = await fetch("http://localhost:3000/status/", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    appeal_id: Number(appeal_id),
                    citizenship_id,
                    voter_card_id,
                    password,
                }),
            });

            const contentType = res.headers.get("content-type") || "";

            if (!res.ok) {
                if (contentType.includes("application/json")) {
                    const err = await res.json();
                    throw new Error(err.message || "Something went wrong");
                } else {
                    const text = await res.text();
                    throw new Error(text || "Something went wrong");
                }
            }

            result = contentType.includes("application/json")
                ? await res.json()
                : null;
        } catch (err) {
            error = err.message;
        } finally {
            loading = false;
        }
    }
</script>

<Navbar />
<div class="h-[80vh] flex justify-center items-center flex-col max-w-xl mx-auto p-6 space-y-4">
    <h2 class="text-6xl font-bold uppercase mb-8">Check Status</h2>

    <div
    class="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80"
    aria-hidden="true"
    >
        <div
            class="relative left-[calc(50%-11rem)] aspect-1155/678 w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-linear-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"
            style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
        ></div>
    </div>
    <input
        class="input input-bordered w-full"
        bind:value={appeal_id}
        placeholder="Appeal ID"
    />
    <input
        class="input input-bordered w-full"
        bind:value={citizenship_id}
        placeholder="Citizenship ID"
    />
    <input
        class="input input-bordered w-full"
        bind:value={voter_card_id}
        placeholder="Voter Card ID"
    />
    <input
        class="input input-bordered w-full"
        type="password"
        bind:value={password}
        placeholder="Password"
    />

    <button
        class="btn btn-primary w-full"
        on:click={checkStatus}
        disabled={loading}
    >
        {#if loading}
            <span class="loading loading-spinner"></span>
        {:else}
            Check Status
        {/if}
    </button>

    {#if error}
        <div class="alert alert-error">
            <span>{error}</span>
        </div>
    {/if}

    {#if result}
        <div class="card shadow-md bg-base-200">
            <div class="card-body">
                <h3 class="card-title">Appeal Found</h3>
                <p><strong>Appeal ID:</strong> {result.appeal_id}</p>
                <p><strong>Status:</strong> {result.status}</p>
                <p><strong>Photos:</strong> {result.photos}</p>
            </div>
        </div>
    {/if}
</div>
<Footer />
