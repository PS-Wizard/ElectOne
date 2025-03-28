const features = [
    { title: "Voter Registration", description: "Register to vote and check your registration status.", icon: "fa-user-check" },
    { title: "Election Information", description: "Get key dates, polling locations, and election rules.", icon: "fa-calendar-check" },
    { title: "Candidate Profiles", description: "View details on candidates, policies, and party affiliations.", icon: "fa-users" },

    { title: "Secure Voting Process", description: "Understand how votes are counted and protected.", icon: "fa-shield-alt" },
    { title: "Results & Updates", description: "Get live election results and statistics.", icon: "fa-chart-line" },
    { title: "Community Engagement", description: "Learn about debates, events, and volunteer opportunities.", icon: "fa-hands-helping" }
];

const container = document.getElementById("features-container");

features.forEach(feature => {
    const card = document.createElement("div");
    card.classList.add("card");
    
    card.innerHTML = `
        <i class="fas ${feature.icon}"></i>
        <h3>${feature.title}</h3>
        <p>${feature.description}</p>
    `;

    container.appendChild(card);
});
