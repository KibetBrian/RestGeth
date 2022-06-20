// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.13;


contract Another{
    address admin;

    //Candidate data structures
    mapping(address=>Vote[]) public candidates;
    Candidate[] candidatesArray;
    uint candidateCounts;

    //Voter data stuctures
    mapping(address=>uint) voters;
    Voter[] votersArray;
    uint votersCount;

    uint count = 0;


    struct Voter{
        address addr;
    }

    struct Candidate{
        string name;
        address addr;
    }
    
    struct Vote {
        address addr;
    }

    constructor(){
        admin = msg.sender;
    }
    modifier onlyAdmin(){
        require(msg.sender==admin, "Only admins are permitted to do this");
        _;
    }

    function CastVote(address _address) public returns (bool) {
        for (uint i=0; i<candidatesArray.length; i++){
            if (candidatesArray[i].addr==_address){
                Vote memory v = Vote({
                    addr: msg.sender
                });
                candidates[_address].push(v);
                return true;
            }
        }
        return false;
    }

    function RegisterCandidate(address _address, string memory _name) onlyAdmin public returns(bool){
        candidates[_address].push();
        candidates[_address].pop();
        
        Candidate memory newCandidate = Candidate({
            name: _name,
            addr: _address
        });
        candidatesArray.push(newCandidate);
        return true;
    }




    function AddVoter (address _address) onlyAdmin public returns (bool){
        voters[_address]=count;
        votersArray.push(Voter(_address));
        count++;
        return true;
    }

    function SeeCandidateVotes(address _address) public view returns (Vote[] memory){
        return candidates[_address];
    }

    function GetCandidates () public view returns (Candidate[] memory){
        return candidatesArray;
    }

    function GetVoters() public view returns(Voter[] memory){
        return votersArray;
    }

    function TotalNumberOfVoter() public view returns(uint){
        return votersArray.length;
    }

}