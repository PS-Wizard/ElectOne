<script>
    import Navbar from "../../../components/Navbar.svelte";

    let news = [
        {
            title: "Election Commission Releases Final Voter List",
            desc: "The updated voter list includes over 2 million new registrants. Officials urge citizens to verify their details before the deadline.",
            img: "https://placehold.co/600x400",
            date: "May 4, 2025",
        },
        {
            title: "Debate Heats Up Between Leading Candidates",
            desc: "A fiery debate between top contenders sparked national conversation on key issues such as inflation and healthcare.",
            img: "https://placehold.co/600x400",
            date: "May 2, 2025",
        },
        {
            title: "Supreme Court To Review Election Disputes",
            desc: "The court will hear appeals from two major districts after claims of voting irregularities and fraud.",
            img: "https://placehold.co/600x400",
            date: "April 30, 2025",
        },
    ];
</script>

<Navbar />

<div class="p-4">
    <h1 class="text-2xl font-bold mb-4 uppercase">Latest Political News</h1>

    <div class="grid md:grid-cols-3 gap-6">
        {#each news as article}
            <div class="card bg-base-100 shadow-xl">
                <figure>
                    <img
                        src={article.img}
                        alt="News Image"
                        class="w-full h-48 object-cover"
                    />
                </figure>
                <div class="card-body">
                    <h2 class="card-title">{article.title}</h2>
                    <p class="text-sm text-gray-500 mb-2">{article.date}</p>
                    <p>{article.desc}</p>
                </div>
            </div>
        {/each}
    </div>
</div>
