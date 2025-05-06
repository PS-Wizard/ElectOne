<script>
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

<div class="max-w-xl mx-auto p-6 space-y-4">
    <h2 class="text-2xl font-bold">Check Appeal Status</h2>

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
