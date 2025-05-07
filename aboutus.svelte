<script>
  // Add any interactivity or props here if needed
</script>

<style>
  * {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
  }
  body {
    font-family: Arial, sans-serif;
    line-height: 1.6;
    background-color: #fff;
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
    content: '\1F539';
    margin-right: 0.5rem;
    color: #007bff;
  }
  footer {
    background-color: #F7F0F0;
    padding: 1rem;
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    align-items: center;
    text-align: center;
    font-size: 0.9rem;
  }
  .footer-left, .footer-center, .footer-right {
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

<main>
  <header>Elect One</header>

  <section class="section">
    <div class="text">
      <h3>🔷 About Us</h3>
      <p>Welcome to ElectOne – Nepal’s trusted Online Voting System. We offer a secure, transparent, and user-friendly platform that modernizes the democratic process. Our mission is to make voting more accessible, ensuring every citizen’s vote is confidential, counted, and respected.</p>
    </div>
    <img src="vote.png" alt="Voting Illustration" />
  </section>

  <section class="section">
    <img src="target.png" alt="Target Illustration" />
    <div class="text">
      <h3>🔷 Our Mission</h3>
      <p>To empower Nepali citizens by providing a reliable and transparent online voting system that ensures every eligible vote counts, from anywhere and at any time.</p>
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
    <img src="features.png" alt="Features Illustration" />
  </section>

  <section class="section">
    <img src="whychoose.png" alt="Security Illustration" />
    <div class="text">
      <h3>🔷 Why Choose Us?</h3>
      <p>We prioritize security, simplicity, and transparency. With features like OTP verification, and role-based dashboards, we ensure that every vote is safe, anonymous, and accurately counted.</p>
    </div>
  </section>

  <footer>
    <div class="footer-left">
      <p><strong>Contact US</strong></p>
      <p>📞 +977-987456321</p>
      <p>📧 electone@gmail.com</p>
    </div>
    <div class="footer-center">
      <p>
        <a href="about.html">About Us</a> |
        <a href="terms.html">Terms & Conditions</a> |
        <a href="cookies.html">Cookie Policy</a>
      </p>
    </div>
    <div class="footer-right">
      <div class="footer-icons">
        <a href="https://www.instagram.com/ElectOne" target="_blank" rel="noopener noreferrer">
          <img src="instagram.png" /> ElectOne
        </a>
        <a href="https://twitter.com/ElectOneNepal" target="_blank" rel="noopener noreferrer">
          <img src="twitter.png" alt="Twitter" /> @ElectOneNepal
        </a>
        <a href="https://www.facebook.com/ElectOneNepal" target="_blank" rel="noopener noreferrer">
          <img src="facebook.png" alt="Facebook" /> Elect One Nepal
        </a>
      </div>
      <div class="footer-logo">
        <img src="Logo.png" alt="Logo" />
      </div>
    </div>
  </footer>
</main>
