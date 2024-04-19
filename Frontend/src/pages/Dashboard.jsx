import { Heading, Container, Box, Button, Text } from "@chakra-ui/react";
import { Link } from "react-router-dom"; // Import Link from react-router-dom
import bubblesImage from "../../public/img/bubbles.png";
import buttercupImage from "../../public/img/buttercup.png";

export default function Dashboard() {
  return (
    <div>
      <Box bgGradient="linear(to-b, white, pink.100, pink)" minHeight="185vh" pt="30px">
        <Container>
          <img src="../../public/img/title.png" alt="Title" style={{ width: '100%', height: 'auto' }}/>
          <img src="../../public/img/ppg.png" alt="PPG" />
        </Container>

        <Box mx="auto" mt="30px" display="flex" alignItems="center" justifyContent="center">
          <img src="../../public/img/boxHome.png" style={{ width: '50%', height: 'auto' }} alt="Home Box" />
        </Box>

        <Heading
          mx="auto"
          mt="100px"
          display="flex"
          alignItems="center"
          justifyContent="center"
          color="white"
          fontFamily={'"Century Gothic", cursive, sans-serif'}
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
