Node findPredecessor(key,n){
    Node pred=n.getPredecessor();
    if (pred==null)
        return n; //n is the current node
    else if (key.isInInterval(pred.ID, n.ID) //check if the key is between the pred and current node
        return n;
    else {
        Node n1=getClosestPrecedingNode(key) //if not,track the closest preceding node and lookup again
        return findPredecessor(key,n1) // recursively  find the predecessor of node n1
    }
}

Node findSuccessor(key,n){
    Node succ=n.getSuccessor();
    if (succ==null)
        return n; //n is the current node
    else if (key.isInInterval(n.ID, succ.ID) //check if the key is between the current node and successor's node
        return n;
    else {
        Node n1=getClosestPrecedingNode(key) //if not,track the closest preceding node and lookup again
        return findSuccessor(key,n1) // recursively findthe predecessor of node n1
    }
}

Set<Serializable> retrieve_R(key){
    hops_R=0; //initialized the hops counter in anti-finger table direction
    while(!retrieved){
        Node responsibleNode_R=null;
        responsibleNode_R = findPredecessor(id);
        hops_R+=1; //while not retrieve the desired key,add the hop counter by 1
        try{
            result_R= responsibleNode_R.retrieveEntries(id); // get the responsibleNode to fetch the entry
            retrieved = true; //if successfully get the value, set retrieved state to true
        }catch(Exception e){

        }
        continue;
    }
    if(result_R !=null) { 
        values1.add(entry.getValue()); //add the lookup result to the valueset
    }
    final_hopsR=hops_R; //get the hop counter for the current lookup operation
    return values1;
}

Set<Serializable> retrieve(key){
    hops=0; //initialized the hops counter in finger table direction
    while(!retrieved){
        Node responsibleNode=null;
        responsibleNode = findSuccessor(id);
        hops+=1; //while not retrieve the desired key, add the hop counter by 1
        try {
            result = responsibleNode.retrieveEntries(id); // get the responsibleNode to fetch the entry
            retrieved = true; //if successfully get the value, set retrieved state to true
        } catch(Exception e){

        }
        continue;
    }
    if(result !=null) {
        values.add(entry.getValue()); // add the lookup result to the valueset
    }
    final_hops=hops; //get the hop counter for the current lookup operation
    return values;
}