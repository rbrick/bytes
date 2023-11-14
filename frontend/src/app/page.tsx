"use client";
import Image from "next/image";
import styles from "./page.module.css";
import next from "next";
import { env } from "process";
import Bytes from "@/components/Bytes";
import { useEffect, useState } from "react";
import { BytesPrice } from "@/types/Bytes";
import { Wallet } from "@/types/Wallet";
import WalletComponent from "@/components/Wallet";
import React from "react";
import {
  Alert,
  AlertHeading,
  Button,
  Col,
  Container,
  FormControl,
  InputGroup,
  Row,
} from "react-bootstrap";
// import getConfig from 'next/config';

// const { publicRuntimeConfig } = getConfig();



function fetchPrice(): Promise<BytesPrice> {
  return new Promise<BytesPrice>((resolve, reject) => {
    fetch(`${process.env.NEXT_PUBLIC_API_URL}/v1/api/price`, {
      method: "GET",
    }).then((response) => {
      response
        .json()
        .then((jsonObj) => {
          resolve(new BytesPrice(jsonObj.etherPrice, jsonObj.usdPrice));
        })
        .catch((err) => {
          console.log(err);
        });
    });
  });
}

function fetchWallet(address: string): Promise<Wallet> {
  return new Promise<Wallet>((resolve, reject) => {
    fetch(`${process.env.NEXT_PUBLIC_API_URL}/v1/api/wallet/${address}`, {
      method: "GET",
    }).then((response) => {
      response
        .json()
        .then((jsonObj) => {
          resolve(new Wallet(address, jsonObj));
        })
        .catch((err) => {
          console.log(err);
        });
    });
  });
}

let ignoreRemount = false;

const ADDRESS_PATTERN = /^0x([a-fA-F0-9]){40}$/;

export default function Home() {
  const [bytesPrice, setBytesPrice] = useState(new BytesPrice("0.00", "0.00"));
  const [displayWallet, setDisplayWallet] = useState(false);
  const [walletAddress, setWalletAddress] = useState("");
  const [error, setError] = useState("");
  const [fetchedWallet, setFetchedWallet] = useState({} as Wallet);

  function onInputChange(text: string) {
    setWalletAddress(text);

    setDisplayWallet(false);
  }

  function onClick() {
    console.log(walletAddress);
    if (!ADDRESS_PATTERN.test(walletAddress)) {
      // console.log(walletAddress)
      setError("invalid address");
      return;
    } else {
      setDisplayWallet(true);
      setError("");
    }
  }

  useEffect(() => {
    if (displayWallet) {
      fetchWallet(walletAddress).then((wallet) => {
        setFetchedWallet(wallet);
      });
    }
  }, [displayWallet, walletAddress]);

  useEffect(() => {
    if (!ignoreRemount) {
      fetchPrice().then((price) => {
        setBytesPrice(price);
      });
    }

    let interval = setInterval(() => {
      fetchPrice().then((price) => {
        setBytesPrice(price);
      });
    }, 5 * 60 * 1000);
    // clear interval on unmount
    return () => {
      ignoreRemount = true;
      clearInterval(interval);
    };
  });

  return (
    <main>
      <Container>
        <Row>
          <h1 style={{ textAlign: "center" }}>bytes.onl</h1>
        </Row>
        <Row>
          <Col sm={3}>
            <Bytes
              priceEther={bytesPrice.formattedEther}
              priceUSD={bytesPrice.formattedUSD}
            />
            &nbsp;&nbsp;
            <span>
              <a href="https://raw.githubusercontent.com/rbrick/bytes/main/CHANGELOG.md">
                Changelog 
              </a>
              &nbsp;
              <a href="#">
                  Roadmap
              </a>


           &nbsp;
           rev-v0.0.1
            &nbsp;
            </span>
          </Col>

          <Col md={6}>
            {error.length > 0 && (
              <Alert dismissible variant="danger" onClose={() => setError("")}>
                {error}
              </Alert>
            )}

            <InputGroup>
              <FormControl
                placeholder="0x..."
                onChange={(event) => onInputChange(event.currentTarget.value)}
              ></FormControl>
              <Button onClick={onClick}>View</Button>
            </InputGroup>
            <br />
            {displayWallet &&
              WalletComponent(fetchedWallet, bytesPrice.priceUSD)}
          </Col>
        </Row>
      </Container>
    </main>
  );
}
