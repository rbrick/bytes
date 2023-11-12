import { Citizen, Wallet } from "@/types/Wallet";
import Image from "next/image";
import {
  Card,
  CardBody,
  CardHeader,
  CardText,
  CardTitle,
  Col,
  Row,
} from "react-bootstrap";

export default function WalletComponent(
  { address, tokenBalance, totalStake, pendingRewards, stakedCitizens }: Wallet,
  priceUSD: number
) {
  let totalStakeNum = Number.parseFloat(totalStake);
  let totalStakeUSDValue = totalStakeNum * priceUSD;
  let pendingRewardsNum = Number.parseFloat(pendingRewards);
  let rewardUSDValue = pendingRewardsNum * priceUSD;
  return (
    <>
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
              {totalStakeNum.toFixed(2)} (${totalStakeUSDValue.toFixed(2)})
            </span>
            <br />
            Pending Rewards{" "}
            <span style={{ fontWeight: "lighter" }}>
              {pendingRewardsNum.toFixed(2)} (${rewardUSDValue.toFixed(2)})
            </span>
          </CardText>

          {stakedCitizens &&
            stakedCitizens.map((citizen) => {
              return (
                <CitizenComponent
                  id={citizen.id}
                  image={citizen.image}
                  season={citizen.season}
                  key={citizen.id}
                />
              );
            })}
        </CardBody>
      </Card>
    </>
  );
}

function CitizenComponent({ image, id, season }: Citizen) {
  return (
    <Card style={{ maxWidth: 200, maxHeight: 150 }}>
      <Row className="g-0">
        <Col xs={4}>
          <Image
            src={image}
            className="img-fluid rounded-start"
            width={100}
            height={100}
            alt=""
          />
        </Col>
        <Col xs={8}>
          <CardBody>
            <CardTitle>ID: {id}</CardTitle>
          </CardBody>
        </Col>
      </Row>
    </Card>
    // <div>
    //     <Image src={image} width={100} height={100} alt=""/>
    //     <span>Season: {season}, ID: {id}</span>
    // </div>
  );
}
