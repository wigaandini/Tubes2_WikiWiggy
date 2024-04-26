import { Box, Container, Flex, Center } from "@chakra-ui/react";
import wImg from "../assets/wiga.png";
import zImg from "../assets/zya.png";
import yImg from "../assets/yudi.png";
import title from "../assets/title.png";
import "../styles/Profile.css";

const Profile = () => {
  return (
    <Box
      bgGradient="linear(to-b, white, #e9dcdc, #cbbdbd)"
      minHeight="90vh"
      display="flex"
      flexDirection="column"
      alignItems="center"
      justifyContent="center"
    >
      
      <Container my={10}>
        <Flex justifyContent="center">
          <Center>
            <img
              src={title}
              alt="Title"
              style={{ width: "100%", height: "auto" }}
            />
          </Center>
        </Flex>
      </Container>

      
      <Container mb={10}>
        <Flex justifyContent="center">
          <Center>
            <img
              src={zImg}
              alt="zya"
              className="shake1-animation"
              style={{ width: "60%", height: "auto", marginRight: "100px" }}
            />

            <img
              src={wImg}
              alt="wiga"
              className="shake2-animation"
              style={{ width: "60%", height: "auto", marginRight: "100px", marginLeft: "100px" }}
            />

            <img
              src={yImg}
              alt="yudi"
              className="shake3-animation"
              style={{ width: "60%", height: "auto", marginLeft: "100px" }}
            />
          </Center>
        </Flex>
      </Container>
    </Box>
  );
};

export default Profile;
