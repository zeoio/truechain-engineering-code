package types

import (
	"strings"
	"fmt"
	"crypto/ecdsa"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
	"time"
)

const (
	CommitteeStart      = iota // start pbft consensus
	CommitteeStop              // stop pbft consensus
	CommitteeSwitchover        //switch pbft committee
	CommitteeOver              // notify current pbft committee end block


	StateUnusedFlag = 0xa0
	StateSwitchingFlag = 0xa1
	StateUsedFlag = 0xa2
	StateRemovedFlag = 0xa3
	StateAddFlag = 0xa4
)

const (
	VoteAgreeAgainst = iota //vote against
	VoteAgree               //vote  agree
)

type CommitteeMembers []*CommitteeMember

type CommitteeMember struct {
	Coinbase  	common.Address
	Publickey 	*ecdsa.PublicKey
	Flag 		int32
}
func (c *CommitteeMember) String() string {
	return fmt.Sprintf("F:%d,C:%s,P:%s",c.Flag,common.ToHex(c.Coinbase[:]),
	common.ToHex(crypto.FromECDSAPub(c.Publickey)))
}

func (g *CommitteeMember) UnmarshalJSON(input []byte) error {
	type committee struct {
		Address common.Address `json:"address,omitempty"`
		PubKey  *hexutil.Bytes `json:"publickey,omitempty"`
		Flag	int32		   `json:"flag,omitempty"`
	}
	var dec committee
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}

	g.Coinbase = dec.Address
	g.Flag = dec.Flag

	var err error
	if dec.PubKey != nil {
		g.Publickey, err = crypto.UnmarshalPubkey(*dec.PubKey)
		if err != nil {
			return err
		}
	}
	return nil
}

type CommitteeNode struct {
	IP        string
	Port      uint
	Port2     uint
	Coinbase  common.Address
	Publickey []byte
}
func (c *CommitteeNode) String() string {
	return fmt.Sprintf("NodeInfo:{IP:%s,P1:%v,P2:%v,Coinbase:%s,P:%s}",c.IP,c.Port,c.Port2,
	common.ToHex(c.Coinbase[:]),common.ToHex(c.Publickey))
}

type PbftSigns []*PbftSign

type PbftSign struct {
	FastHeight *big.Int
	FastHash   common.Hash // fastblock hash
	Result     uint        // 0--agree,1--against
	Sign       []byte      // sign for fastblock height + hash + result
}

type PbftAgentProxy interface {
	FetchFastBlock(committeeId *big.Int,infos *SwitchInfos) (*Block, error)
	VerifyFastBlock(*Block) (*PbftSign, error)
	BroadcastFastBlock(*Block)
	BroadcastConsensus(block *Block) error
	GetCurrentHeight() *big.Int
}

type PbftServerProxy interface {
	PutCommittee(committeeInfo *CommitteeInfo) error
	UpdateCommittee(info *CommitteeInfo) error
	PutNodes(id *big.Int, nodes []*CommitteeNode) error
	Notify(id *big.Int, action int) error
	SetCommitteeStop(committeeId *big.Int, stop uint64) error
	GetCommitteeStatus(committeeID *big.Int) map[string]interface{}
}

// Hash returns the block hash of the PbftSign, which is simply the keccak256 hash of its
// RLP encoding.
func (h *PbftSign) Hash() common.Hash {
	return rlpHash(h)
}

func (h *PbftSign) HashWithNoSign() common.Hash {
	return rlpHash([]interface{}{
		h.FastHeight,
		h.FastHash,
		h.Result,
	})
}

type CommitteeInfo struct {
	Id      		*big.Int
	StartHeight		*big.Int
	Members 		[]*CommitteeMember
	BackMembers		[]*CommitteeMember
}
func (c *CommitteeInfo) String() string{
	if c.Members != nil {
		memStrings := make([]string, len(c.Members))
		for i, m := range c.Members {
			if m == nil {
				memStrings[i] = "nil-Member"
			} else {
				memStrings[i] = m.String()
			}
		}
		return fmt.Sprintf("CommitteeInfo{ID:%s,SH:%s,M:{%s}}",c.Id,c.StartHeight,strings.Join(memStrings,"\n  "))
	}
	return fmt.Sprintf("CommitteeInfo{ID:%s,SH:%s}",c.Id,c.StartHeight)
}

type EncryptCommitteeNode []byte
type Sign []byte

type EncryptNodeMessage struct {
	CreatedAt   time.Time
	CommitteeId *big.Int
	Nodes       []EncryptCommitteeNode
	Sign        //sign msg
}

func (c *EncryptNodeMessage) HashWithoutSign() common.Hash {
	return RlpHash([]interface{}{
		c.Nodes,
		c.CommitteeId,
	})
}

func (c *EncryptNodeMessage) Hash() common.Hash {
	return RlpHash(c)
}

func RlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

const (
	SwitchAppend = 0xEE
	SwitchRemove = 0xAA
)

type SwitchEnter struct {
	Pk  		[]byte
	Flag 		int32
}
func (s *SwitchEnter) String() string {
	return fmt.Sprintf("p:%s,s:%d",common.ToHex(s.Pk),s.Flag)
}

type SwitchInfos struct {
	CID 	uint64
	Vals 	[]*SwitchEnter
}

func (s *SwitchInfos) String() string {
	memStrings := make([]string, len(s.Vals))
	for i, m := range s.Vals {
		if m == nil {
			memStrings[i] = "nil-Member"
		} else {
			memStrings[i] = m.String()
		}
	}
	return fmt.Sprintf("SwitchInfos{CID:%d,Vals:{%s}}",s.CID,strings.Join(memStrings,"\n  "))
}