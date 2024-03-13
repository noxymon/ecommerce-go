/*!
* Start Bootstrap - Shop Homepage v5.0.6 (https://startbootstrap.com/template/shop-homepage)
* Copyright 2013-2023 Start Bootstrap
* Licensed under MIT (https://github.com/StartBootstrap/startbootstrap-shop-homepage/blob/master/LICENSE)
*/
// This file is intentionally blank
// Use this file to add JavaScript to your project
$(document).ready(function () {
  prepareSearch();
});

function prepareSearch() {
  console.log("search start !!")
  const searchClient = algoliasearch('NIFL4FEXTO', 'c40abc3344706c260d2e150e59f6586d');

  const search = instantsearch({
    indexName: 'demo_ecommerce',
    searchClient,
  });

  search.addWidgets([
    instantsearch.widgets.searchBox({
      container: '#searchbox',
    }),

    instantsearch.widgets.hits({
      container: '#hits',
    })
  ]);

  search.start();
}

