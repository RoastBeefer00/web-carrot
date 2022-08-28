<script>
    function removeRecipe(recipe) {
        for (let index = 0; index < $storeFE.length; index++) {
            if ($storeFE[index] == recipe) {
                var l = $undo.length;
                $undo[l] = {
                    "task": "Delete",
                    "index": index,
                    "recipe": recipe
                }
            } 
        }
        
        console.log($undo)
		$storeFE = $storeFE.filter(r => r !== recipe);
	}

    async function replaceRecipe(recipe) {
        for (let index = 0; index < $storeFE.length; index++) {
            if ($storeFE[index] == recipe) {
                var l = $undo.length;
                $undo[l] = {
                    "task": "Replace",
                    "index": index,
                    "recipe": recipe
                }
                let response = await fetch("https://r7qi88.deta.dev/random");
		        let recipes = await response.json();
                console.log(recipes[0]);
                $storeFE[index] = recipes[0];
            } 
        }
        console.log($undo)
    }

    function toggle() {
        visible = !visible;
    }

    export let recipes;
    let visible = false;

	import {
		Card,
		CardBody,
		CardHeader,
		CardSubtitle,
		CardTitle,
		Input,
		ListGroup,
		ListGroupItem,
		Container,
		Row,
		Col,
        Icon
	} from 'sveltestrap';

    import { storeFE, undo } from '../scripts/store.js'
    import { slide } from 'svelte/transition';
</script>

<div>
    <Card class="mb-3">
        {#await recipes}
            <p>Loading...</p>
        {:then recipes}
            <div class="header">
                <h3>{recipes.title}</h3>
                <p class="time">{recipes.time}</p>
                <button class="button button_normal" on:click={toggle}><Icon name={visible ? "chevron-up" : "chevron-down"} /></button>
                <button class="button button_normal button_right" on:click={removeRecipe(recipes)}><Icon name="trash" /></button>
                <button class="button button_normal button_right" on:click={replaceRecipe(recipes)}><Icon name="arrow-repeat" /></button>
            </div>
            {#if visible}
            <div transition:slide>
                <CardBody style="background:#8EE4AF">
                    <CardTitle style="color:#05386B">Ingredients</CardTitle>
                    <Container>
                        <Row cols={3}>
                            {#each recipes.ingredients as ingredient}
                                <Col>
                                    <Input style="color:#05386B" id="c1" type ="checkbox" label={ingredient} />
                                </Col>
                            {/each}
                        </Row>
                    </Container>
                    <CardTitle style="color:#05386B">Steps</CardTitle>
                    <ListGroup numbered>
                    {#each recipes.steps as step}
                        <ListGroupItem style="background:#EDF5E1; color:#05386B">{step}</ListGroupItem>
                    {/each}
                    </ListGroup>
                </CardBody>
            </div>
            {/if}
        {/await}
    </Card>
</div>

<style>
    .header {
        background: #05386B;
    }
    
    h3 {
        color: #EDF5E1;
        font-size: 16pt;
        margin: 5px;
    }

    .time {
        color: #EDF5E1;
        margin: 5px;
    }

    .button_normal {
        background: #379683; 
        color: #EDF5E1;
    }
</style>
