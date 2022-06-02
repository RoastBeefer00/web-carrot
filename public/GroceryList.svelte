<script>
  import { storeFE } from '../scripts/store.js'

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
  import { beforeUpdate } from 'svelte';

  let open = false;
  const toggle = () => (open = !open);

  // const ingredientsList = [];

  // afterUpdate(() => {
  //   ingredientsList = [];
  //   for (let index = 0; index < $storeFE.length; index++) {
  //     $storeFE.ingredients.forEach(ingredient => {
  //       ingredientsList.push(ingredient);
  //     });
  //   }
  //   // [a-zA-Z].*
  //   ingredientsList = ingredientsList.sort();
  // });
</script>

<div>
  <Button style="background:#05386B; color:#EDF5E1" on:click={toggle} disabled={$storeFE.length == 0}>Grocery List</Button>
  <Modal isOpen={open} {toggle}>
    <ModalHeader {toggle} style="background:#379683; color:#EDF5E1">Grocery List</ModalHeader>
    <ModalBody style="background:#EDF5E1">
      <Container>
        <Row cols={2} style="max-height:calc(100vh - 225px); overflow:scroll">
          <div >
          {#each $storeFE as recipe}
            {#each recipe.ingredients.sort() as ingredient}
              <Input style="color:#05386B" id="c1" type ="checkbox" label={ingredient} />
            {/each}
          {/each}
          </div>
        </Row>
      </Container>
    </ModalBody>
    <ModalFooter style="background:#379683">
      <Button color="secondary" on:click={toggle} style="background:#05386B">Close</Button>
    </ModalFooter>
  </Modal>
</div>