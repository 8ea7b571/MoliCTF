function prevPage() {
    const currentPage = document.getElementById('some-data').getAttribute('data-current-page');
    if (currentPage > 1) {
        const prevPage = parseInt(currentPage) - 1;
        window.location.href = '/teams?page=' + prevPage;
    }
}

function nextPage() {
    const currentPage = document.getElementById('some-data').getAttribute('data-current-page');
    const totalPages = document.getElementById('some-data').getAttribute('data-total-page');

    if (currentPage < totalPages) {
        const nextPage = parseInt(currentPage) + 1;
        window.location.href = '/teams?page=' + nextPage;
    }
}