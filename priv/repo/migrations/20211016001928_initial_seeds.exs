defmodule Repo.Migrations.InitialSeeds do
  use Ecto.Migration

  def up do
    # set up master application entity
    DB.execute("""
      INSERT INTO application_entity (
        pub_id,
        external_id,
        type
      )
      VALUES (
        'MASTER',
        'MASTER',
        'MASTER'
      );
    """)
  end

  def down do

  end
end
