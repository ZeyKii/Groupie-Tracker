const filters = document.getElementById('filterBox');
const button = document.getElementById('filterButton'); 


button.onclick = function () {
    if (filters.style.display === 'block') {
        filters.style.display = 'none';
        button.innerHTML = 'Show Filters';

    } else {
        filters.style.display = 'block';
        button.innerHTML = 'Hide Filters';
    };
};

