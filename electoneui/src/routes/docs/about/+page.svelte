<script>
    import Footer from "../../../components/footer.svelte";
    import Navbar from "../../../components/Navbar.svelte";

    // Add any interactivity or props here if needed
</script>

<Navbar />
<main>
    <header>Elect One</header>
    <section class="section">
        <div class="text">
            <h3>🔷 About Us</h3>
            <p>
                Welcome to ElectOne – Nepal’s trusted Online Voting System. We
                offer a secure, transparent, and user-friendly platform that
                modernizes the democratic process. Our mission is to make voting
                more accessible, ensuring every citizen’s vote is confidential,
                counted, and respected.
            </p>
        </div>
        <img src="/vote.png" alt="Voting Illustration" />
    </section>

    <section class="section">
        <img src="/target.png" alt="Target Illustration" />
        <div class="text">
            <h3>🔷 Our Mission</h3>
            <p>
                To empower Nepali citizens by providing a reliable and
                transparent online voting system that ensures every eligible
                vote counts, from anywhere and at any time.
            </p>
        </div>
    </section>

    <section class="section features">
        <div class="text">
            <h3>🔷 Key Features</h3>
            <ul>
                <li>🔐 Secure Login & Voter Authentication</li>
                <li>🗳️ One Person, One Vote Enforcement</li>
                <li>👥 Verified Candidate & Voter Profiles</li>
                <li>📊 Live Election Results</li>
                <li>📄 Access to Election Information</li>
                <li>👰️ Admin Controls to Manage Elections</li>
            </ul>
        </div>
        <img src="/features.png" alt="Features Illustration" />
    </section>

    <section class="section">
        <img src="/whychoose.png" alt="Security Illustration" />
        <div class="text">
            <h3>🔷 Why Choose Us?</h3>
            <p>
                We prioritize security, simplicity, and transparency. With
                features like OTP verification, and role-based dashboards, we
                ensure that every vote is safe, anonymous, and accurately
                counted.
            </p>
        </div>
    </section>
    <div
        class="absolute inset-x-0 top-[calc(100%-13rem)] -z-10 transform-gpu overflow-hidden blur-3xl sm:top-[calc(100%-30rem)]"
        aria-hidden="true"
    >
        <div
            class="relative left-[calc(50%+3rem)] aspect-1155/678 w-[36.125rem] -translate-x-1/2 bg-linear-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%+36rem)] sm:w-[72.1875rem]"
            style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
        ></div>
    </div>
    <Footer />
</main>

<style>
    * {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
    }
    body {
        font-family: Arial, sans-serif;
        line-height: 1.6;
        color: #333;
    }
    header {
        text-align: center;
        padding: 1rem;
        font-size: 2rem;
        font-weight: bold;
    }
    .section {
        padding: 1.5rem;
        display: flex;
        flex-wrap: wrap;
        align-items: center;
        justify-content: center;
        gap: 1.5rem;
    }
    .section img {
        max-width: 300px;
        width: 100%;
    }
    .text {
        flex: 1 1 300px;
        max-width: 600px;
    }
    .features {
        padding: 1.5rem;
    }
    .features ul {
        list-style: none;
        padding-left: 1rem;
    }
    .features li::before {
        content: "\1F539";
        margin-right: 0.5rem;
        color: #007bff;
    }
    footer {
        padding: 1rem;
        display: flex;
        flex-wrap: wrap;
        justify-content: space-between;
        align-items: center;
        text-align: center;
        font-size: 0.9rem;
    }
    .footer-left,
    .footer-center,
    .footer-right {
        flex: 1 1 200px;
        margin: 0.5rem;
    }
    .footer-icons {
        margin-top: 0.5rem;
    }
    .footer-icons a {
        display: inline-block;
        margin: 0 0.3rem;
        text-decoration: none;
        color: inherit;
    }
    .footer-icons img {
        height: 24px;
        margin-right: 0.3rem;
        vertical-align: middle;
    }
    .footer-logo {
        margin-top: 1rem;
        text-align: right;
    }
    .footer-logo img {
        width: 100px;
        height: auto;
        max-height: 80px;
        object-fit: contain;
    }
    @media (max-width: 768px) {
        .section {
            flex-direction: column;
            text-align: center;
        }
        footer {
            flex-direction: column;
        }
        .footer-logo {
            text-align: center;
            margin-top: 1rem;
        }
    }
</style>
