<script>
  // Optional: you can add props or logic here if needed in the future
</script>

<style>
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
    footer {
      flex-direction: column;
    }
    .footer-logo {
      text-align: center;
      margin-top: 1rem;
    }
  }
</style>

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
        <img src="instagram.png" alt="Instagram" /> ElectOne
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
