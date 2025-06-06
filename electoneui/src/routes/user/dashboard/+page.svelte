<script>
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    import Navbar from "../../../components/Navbar.svelte";
    import Footer from "../../../components/footer.svelte";

    let elections = [];
    let loading = true;
    let error = null;

    let newsArticles = [];

    function hasElectionStarted(startDate) {
        return new Date(startDate) <= new Date();
    }

    onMount(async () => {
        const response = await fetch(
            "https://newsapi.org/v2/top-headlines?country=us&pageSize=6&apiKey=499e05a3650a4d31bb11d2c9f4ec5f74",
        );
        const data = await response.json();
        newsArticles = data.articles;
    });

    function hasElectionEnded(endDate) {
        return new Date(endDate) < new Date();
    }
    onMount(async () => {
        const token = localStorage.getItem("user_token");
        if (!token) {
            error = "You're not logged in.";
            loading = false;
            return;
        }

        try {
            const res = await fetch(
                "http://localhost:3000/election?limit=10&offset=0",
                {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to fetch elections");
            const data = await res.json();
            elections = data ?? []; // make sure it's at least an empty array
        } catch (err) {
            error = err.message;
        } finally {
            loading = false;
        }
    });
</script>

<Navbar />
<section class="lg:border-2 lg:border-gray-300 rounded-lg m-4">
    <section class="m-2 lg:m-12">
        <div
            class="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80"
            aria-hidden="true"
        >
            <div
                class="relative left-[calc(50%-11rem)] aspect-1155/678 w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-linear-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"
                style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
            ></div>
        </div>
        <h1 class="text-2xl font-bold mb-4 uppercase tracking-wide">
            Elections
        </h1>

        {#if loading}
            <p class="text-gray-500">Loading elections...</p>
        {:else if error}
            <p class="text-red-500">Error: {error}</p>
        {:else if elections.length === 0}
            <p class="text-gray-500">There are no elections</p>
        {:else}
            <div
                class="grid md:grid-cols-2 lg:grid-cols-3 gap-6 cursor-pointer"
            >
                {#each elections as e}
                    {#if hasElectionStarted(e.start_date)}
                        <div
                            class="card bg-base-100 border-2 border-gray-400 transition-transform hover:scale-105"
                            on:click={() =>
                                !hasElectionEnded(e.end_date) &&
                                goto(
                                    `/user/dashboard/election/${e.election_id}`,
                                )}
                        >
                            <div class="card-body">
                                <h2 class="card-title">{e.name}</h2>
                                <p class="text-sm text-gray-500">
                                    {e.location}
                                </p>
                                <p class="text-sm text-gray-600">
                                    {e.description}
                                </p>
                                <div
                                    class="mt-2 flex flex-wrap gap-2 items-center"
                                >
                                    <span class="badge badge-outline"
                                        >Start: {e.start_date}</span
                                    >
                                    <span class="badge badge-outline"
                                        >End: {e.end_date}</span
                                    >

                                    {#if hasElectionEnded(e.end_date)}
                                        <a
                                            href={`/results/${e.election_id}/`}
                                            class="badge bg-blue-500 text-white hover:bg-blue-600"
                                            on:click|stopPropagation
                                        >
                                            View Election Analytics
                                        </a>
                                    {/if}
                                </div>
                            </div>
                        </div>
                    {/if}
                {/each}
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
    <hr class="my-16 border-gray-300" />
    <section class="mt-12">
        <h2 class="text-2xl font-bold m-12 uppercase tracking-wide">
            Latest News
        </h2>
        <div class="m-2 lg:m-12 grid md:grid-cols-2 lg:grid-cols-3 gap-6">
            {#each newsArticles as article}
                <div
                    class="card bg-base-100 border border-gray-200 p-4 rounded-lg transition-transform hover:scale-105"
                >
                    <img
                        src={article.urlToImage}
                        alt={article.title}
                        class="w-full h-48 object-cover rounded-md"
                    />
                    <h3 class="mt-2 text-lg font-semibold">{article.title}</h3>
                    <p class="text-sm text-gray-600">{article.description}</p>
                    <a
                        href={article.url}
                        target="_blank"
                        class="text-blue-500 mt-2 inline-block">Read more</a
                    >
                </div>
            {/each}
        </div>
    </section>
</section>
<Footer />
