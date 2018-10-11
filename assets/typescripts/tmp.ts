if (manualAll) {
  let manualUpdate = new MatchResult.Result[];
  manualAll.forEach(ori => {
    const exist = modifyResults.find(ori=>ori.SetTypeID === mod.SetTypeID && ori.BoxscoreTypeID === mod.BoxscoreTypeID)
    if (!exist) {
      manualUpdate.push(ori);
    }
  });
  manualUpdate.push(...modifyResults);
  this.manualMaps.set(matchID, manualUpdate);
}

/////////////
if (modifyResults.length > 0) {
  const matchID = this.data.matchID;
  const manualAll = this.manualMaps.get(matchID);
  // console.log(manualAll);
  // if (manualAll) {
  //   manualAll.push(...modifyResults);
  //   this.manualMaps.set(matchID, manualAll);
  // } else {
  //   this.manualMaps.set(matchID, modifyResults);
  // }
  this.manualMaps.set(matchID, modifyResults);
  this.store.dispatch(new ScoreUpdate(modifyResults));
  if (this.rememberCheck === 5) {
    this.store.dispatch(new SelectResultSource(this.rememberCheck, this.FTResult));
  } else if (this.rememberCheck === 6) {
    this.store.dispatch(new SelectResultSource(this.rememberCheck, this.OtherResult));
  }

}