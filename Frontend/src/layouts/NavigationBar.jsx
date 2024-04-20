import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { Image, Text, Box, Flex, Stack } from "@chakra-ui/react";
import { MdClose, MdMenu } from "react-icons/md";
import logo from "../assets/3.png";
import MobileNavigation from "./MobileNavigation";

const Logo = () => {
  return (
    <Link to="/">
      <Box
        display={"flex"}
        flexDirection={"row"}
        justifyContent="start"
        alignItems="center"
        _hover={{ color: "#e29ccc" }}
      >
        <Image src={logo} w={"50px"} mr={{ base: "30px" }} style={{ width: '14%', height: 'auto' }} />
        <Text fontFamily="monospace" fontWeight="bold" fontSize="3xl">
          WikiRace
        </Text>
      </Box>
    </Link>
  );
};

const MenuItem = ({ children, isLast, to = "/", ...rest }) => {
  return (
    <Link to={to} _hover={{ bg: "#DFDF96" }}>
      <Box
        paddingX={{ md:"0.5em", lg:"1.5em"}}
        paddingY={"2em"}
        display="flex"
        alignItems={"center"}
        justifyContent={"center"}
        {...rest}
        borderBottomWidth={"5px"}
        borderColor={"transparent"}
        _hover={{ 
            color: "#e29ccc", 
            borderColor: "#e29ccc",
            transitionDuration:"0.25s",
            transitionTimingFunction:"ease-in-out"}}
      >
        <Text fontFamily="monospace" fontWeight="bold" fontSize="20px">{children}</Text>
      </Box>
    </Link>
  );
};

const NavToggle = ({ toggle, isOpen }) => {
  return (
    <Box display={{ base: "block", md: "none" }} onClick={toggle}>
      {isOpen ? <MdClose size={24} /> : <MdMenu size={24} />}
    </Box>
  );
};

const MenuContainer = () => {
  return (
    <Box
      display={{ base: "none", md: "block" }}
      flexBasis={{ base: "100%", md: "auto" }}
      py={2}
    >
      <Stack
        spacing={0.5}
        align="center"
        justify={["center", "space-between", "flex-end", "flex-end"]}
        direction={["column", "row", "row", "row"]}
        pt={[6, 6, 0, 0]}
      >
        <MenuItem to="" >Home</MenuItem>
        <MenuItem to="profile" >About Us</MenuItem>
        <MenuItem to="how-to-use">How To Use</MenuItem>
      </Stack>
    </Box>
  );
};

function NavigationBar() {
  const [menuOpen, setMenuOpen] = useState(false);
  const toggle = () => setMenuOpen(!menuOpen);
  const [navbarPos, setNavbarPos] = useState(0);

  useEffect(() => {
      var prevScrollPosY = window.pageYOffset

      const detectScrollY = () => {        
        var temp = window.scrollY
        if (temp > prevScrollPosY) {
          setNavbarPos(0)
        } else {
          setNavbarPos(-32)
        }
        prevScrollPosY = temp
      }

      window.addEventListener("scroll", detectScrollY)

      return () => {
        window.removeEventListener("scroll", detectScrollY)
      }
    }
  )

  return (
    <>
        <Flex
        bg={"white"}
        as="nav"
        position={"sticky"}
        top={navbarPos}
        insetX={0}
        align="center"
        justify="space-between"
        wrap="wrap"
        px={{base:5, sm: 10}}
        py={{base:5, md:0}}
        shadow={"md"}
        zIndex={2}
        transitionDuration={"0.3s"}
        transitionTimingFunction={"ease-in-out"}>
            <Logo />
            <MobileNavigation
            open={menuOpen}
            closeMenu={() => setMenuOpen(false)}
            />
            <NavToggle isOpen={menuOpen} toggle={toggle} />
            <MenuContainer />
        </Flex>
    </>
  );
}

export default NavigationBar;