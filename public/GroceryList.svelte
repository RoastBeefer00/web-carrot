<script>
  import { storeFE, groceryList } from '../scripts/store.js'
  import { derived } from 'svelte/store';

  import {
    Button,
    Modal,
    ModalHeader,
    ModalFooter,
    ModalBody,
    Input,
    Container,
    Row
  } from 'sveltestrap'

  let open = false;
  // const toggle = () => (open = !open);
  function toggle() {
    open = !open;
    splitIngredients();
  }

  $: sortedGroceryList = sortByIngredient(Object.values($groceryList));

  function sortByIngredient(array) {
    array.sort((el1, el2) =>
      el1.ingredient.toLowerCase().localeCompare(el2.ingredient.toLowerCase())
    );
    combineIngredients(array);
    return array;
  }

  function splitIngredients() {
    $groceryList = [];
    var value;
    var rest;
    var json;
    $storeFE.forEach(recipe => {
      recipe.ingredients.forEach(ingredient => {
      value = ingredient.match("^\\d*[^a-zA-Z \\*]?\\d*|¼|¾|½|⅓");
      //console.log(value);

      rest = ingredient.match("\\*?[a-zA-Z].*")
      //console.log(rest);

      json = "{\"quantity\":\"" + value + "\", \"ingredient\":\"" + rest + "\"}";
      $groceryList[$groceryList.length] = JSON.parse(json);
      });
    });
  }

  function combineIngredients(array) {
    for (let i = 0; i < array.length; i++) {
      for (let j = i+1; j < array.length; j++) {
        if (array[i].ingredient.includes(array[j].ingredient) || array[j].ingredient.includes(array[i].ingredient)) {
          if (array[i].quantity != "") {
            console.log(parseFloat(array[i].quantity) + parseFloat(array[j].quantity));
          }
          console.log(array[i].ingredient);
        }
      } 
    }
    console.log(parseFloat(".5") + 1);
  }

  function printIngredient(ingredient) {
    if (ingredient.quantity == '') {
      return ingredient.ingredient;
    }
    return ingredient.quantity + " " + ingredient.ingredient
  }
</script>

<div>
  <Button style="background:#05386B; color:#EDF5E1" on:click={toggle} disabled={$storeFE.length == 0}>Grocery List</Button>
  <Modal isOpen={open} {toggle}>
    <ModalHeader {toggle} style="background:#379683; color:#EDF5E1">Grocery List</ModalHeader>
    <ModalBody style="background:#EDF5E1">
      <Container>
        <Row cols={1} style="max-height:calc(100vh - 300px); overflow:scroll">
          {#each sortedGroceryList as ingredient}
              <Input style="color:#05386B" id="c1" type ="checkbox" label={printIngredient(ingredient)} />

          {/each}
        </Row>
      </Container>
    </ModalBody>
    <ModalFooter style="background:#379683">
      <!-- <Button color="secondary" on:click={splitIngredients} style="background:#05386B">Close</Button> -->
    </ModalFooter>
  </Modal>
</div>