import Sequelize, { Model } from 'sequelize';

class Social extends Model {
  static init(sequelize) {
    super.init(
      {
        date: Sequelize.DATE,
        name_person: Sequelize.STRING,
        address: Sequelize.STRING,
        type: Sequelize.STRING,
        age: Sequelize.STRING,
        bible: Sequelize.STRING,
        evangelical: Sequelize.STRING,
        contact: Sequelize.STRING,
        frequents_church: Sequelize.STRING,
        cult_home: Sequelize.STRING,
        study_bible: Sequelize.STRING,
        study_confimation: Sequelize.STRING,
        prayer_request: Sequelize.STRING,
        reconciled: Sequelize.STRING,
        visit: Sequelize.STRING,
        accept_christ: Sequelize.STRING,
        medical: Sequelize.STRING,
        optician: Sequelize.STRING,
        pressure: Sequelize.STRING,
        glucose: Sequelize.STRING,
        aesthetics: Sequelize.STRING,
        cutting_hair: Sequelize.STRING,
        hairstyle: Sequelize.STRING,
        Nail: Sequelize.STRING,
        Eyebrow: Sequelize.STRING,
        note: Sequelize.STRING,
      },
      { sequelize }
    );

    return this;
  }
}

export default Social;
