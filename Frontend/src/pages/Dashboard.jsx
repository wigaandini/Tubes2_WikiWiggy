import { Heading, Container, Box, Button } from "@chakra-ui/react";
import { Link } from "react-router-dom";
import bubblesImage from "../assets/bubbles.png";
import buttercupImage from "../assets/buttercup.png";

export default function Dashboard() {
  return (
    <div>
      <Box bgGradient="linear(to-b, white, #fbd7e6, #db9fb8)" minHeight="185vh" pt="30px">
        <Container>
          <img src="./img/title.png" alt="Title" style={{ width: '100%', height: 'auto' }}/>
          <img src="./img/ppg.png" alt="PPG" />
        </Container>

        <Box mx="auto" mt="30px" display="flex" alignItems="center" justifyContent="center">
          <img src="./img/boxHome.png" style={{ width: '50%', height: 'auto' }} alt="Home Box" />
        </Box>

        <Heading
          mx="auto"
          mt="100px"
          display="flex"
          alignItems="center"
          justifyContent="center"
          color="white"
          fontFamily="monospace"
          textShadow="2px 2px 4px rgba(0,0,0,0.4)"
        >
          How does it work??
        </Heading>

        <Box mx="auto" mt="140px" display="flex" alignItems="center" justifyContent="center">

          <Box textAlign="center">
            <Button as={Link} to="/bfs-page" variant="none" _hover={{ bg: "transparent", transform: "scale(1.1)", transition: "transform 0.3s ease-in-out" }}>
              <img src={bubblesImage} alt="Bubbles" style={{ width: '50%', height: 'auto' }} />
            </Button>
          </Box>

          <Box textAlign="center">
            <Button as={Link} to="/ids-page" variant="none" _hover={{ bg: "transparent", transform: "scale(1.1)", transition: "transform 0.3s ease-in-out"  }}>
              <img src={buttercupImage} alt="Buttercup" style={{ width: '50%', height: 'auto' }} />
            </Button>
          </Box>

        </Box>

      </Box>
    </div>
  );
}
