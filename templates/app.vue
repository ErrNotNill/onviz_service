<script src="https://cdn.jsdelivr.net/npm/vue@latest"></script>

<script>

created() {
// GET request using fetch with set headers
const headers = { "Content-Type": "application/json" };
fetch("https://api.npms.io/v2/search?q=vue", { headers })
.then(response => response.json())
.then(data => (this.totalVuePackages = data.total));
}

</script>