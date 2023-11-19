import { Citizen, Wallet } from "@/types/Wallet";
import Image from "next/image";
import React, { Component, ReactNode } from "react";
import {
  Card,
  CardBody,
  CardHeader,
  CardText,
  CardTitle,
  Col,
  Container,
  Row,
} from "react-bootstrap";




let currencyFormatter = new Intl.NumberFormat('en-US', {currency: 'USD', style: 'currency'});


export default function WalletComponent(
  { address, tokenBalance, totalStake, pendingRewards, stakedCitizens }: Wallet,
  priceUSD: number
) {
  let totalStakeNum = Number.parseFloat(totalStake);
  let totalStakeUSDValue = totalStakeNum * priceUSD;
  let pendingRewardsNum = Number.parseFloat(pendingRewards);
  let rewardUSDValue = pendingRewardsNum * priceUSD;
  return (
    <Container fluid>
      <Card>
        <CardHeader>Summary</CardHeader>
        <CardBody>
          <CardText>
            Balance:{" "}
            <span style={{ fontWeight: "lighter" }}>
              {Math.round(Number.parseFloat(tokenBalance)).toFixed(2)}
            </span>
            <br />
            Total Stake:{" "}
            <span style={{ fontWeight: "lighter" }}>
              {totalStakeNum.toFixed(2)} ({currencyFormatter.format(totalStakeUSDValue)})
            </span>
            <br />
            Pending Rewards{" "}
            <span style={{ fontWeight: "lighter" }}>
              {pendingRewardsNum.toFixed(2)} ({currencyFormatter.format(rewardUSDValue)})
            </span>
          </CardText>

          <CardTitle>Staked Citizens</CardTitle>
          {stakedCitizens && CitizenComponent(stakedCitizens)}
        </CardBody>
      </Card>
    </Container>
  );
}

function CitizenComponent(citizens: Citizen[]) {
  let array = [];

  for (let i = 0; i < (citizens.length / 2) + citizens.length % 2; i++) {
   let subArray = citizens.slice((i * 2), (i + 1) * 2);

   console.log(subArray.length);
   let cards = subArray.map((citizen, index) => {
      return (
        <Card style={{ maxWidth: 200, maxHeight: 150, margin: 2 }} key={citizen.id}>
        <Row className="g-0">
          <Col xs={4}>
            <Image
              src={citizen.image}
              className="img-fluid rounded-start"
              width={100}
              height={100}
              alt=""
            />
          </Col>
          <Col xs={8}>
            <CardBody>
              <CardTitle>ID: {citizen.id}</CardTitle>
            </CardBody>
          </Col>
        </Row>
      </Card>
      )
    });


    console.log("card length", cards.length)


    array.push( (
      <>
      
        <Row className="g-0 no-gutters">
        {cards}
        </Row>
      </>
       
    ) );
  }

  return array
}
