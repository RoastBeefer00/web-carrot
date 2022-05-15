<script>
    export let recipes;

    function removeComponent(title) {
		$storeFE = $storeFE.filter(r => r.title !== title);
	}

    import { storeFE } from '../scripts/store.js'

	import { 
		Button,
		Card,
		CardBody,
		CardFooter,
		CardHeader,
		CardSubtitle,
		CardText,
		CardTitle,
		FormGroup,
		Input,
		ListGroup,
		ListGroupItem,
		Container,
		Row,
		Col
	} from 'sveltestrap';
</script>

<div>
    <Card class="mb-3">
        {#await recipes}
            <p>Loading...</p>
        {:then recipes}
            <CardHeader>
                <CardTitle>{recipes.title}</CardTitle>
                <CardSubtitle>Prep Time: {recipes.time}</CardSubtitle>
                <Button on:click={removeComponent(recipes.title)}>Remove</Button>
            </CardHeader>
            <CardBody>
                <CardTitle>Ingredients</CardTitle>
                <Container>
                    <Row cols={4}>
                        {#each recipes.ingredients as ingredient}
                            <Col>
                                <Input id="c1" type ="checkbox" label={ingredient} />
                            </Col>
                        {/each}
                    </Row>
                </Container>
                <h1>Steps</h1>
                <ListGroup numbered>
                {#each recipes.steps as step}
                    <ListGroupItem>{step}</ListGroupItem>
                {/each}
                </ListGroup>
            </CardBody>
        {:catch error}
            <p style="color: red">{error.message}</p>
        {/await}
    </Card>
</div>