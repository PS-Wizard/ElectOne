<script lang="ts">
    function alertCookies() {
        alert(document.cookie);
    }

    function setRandomCookie() {
        const randomValue = Math.random().toString(36).substring(2, 15);
        document.cookie = `randomCookie=${randomValue}; path=/; SameSite=None; Secure=false`;
        alert(`Set randomCookie to: ${randomValue}`);
    }
</script>

<button on:click={alertCookies}>Alert Cookies</button>
<button on:click={setRandomCookie}>Set Random Cookie</button>
<button on:click={alertCookies}>Alert Cookies Again</button>
