CREATE TABLE createdTokens(id serial PRIMARY KEY, id_addr_cntr_TokenTRI varchar(50), id_addr_cntr_Auction varchar(50));
CREATE TABLE tokenoffchain(id serial PRIMARY KEY, log_id varchar(50), last_block_stored varchar(50), timestamp varchar(50), token_owner varchar(50), token_id varchar(50), token_name varchar(50), token_ler varchar(50), token_state varchar(50), token_ttl varchar(50), token_block_created varchar(50), batch_id varchar(50), batch_attr_gen varchar(50), batch_location varchar(50), participant_id varchar(50), participant_name varchar(50), participant_contact varchar(50), participant_roleProductor varchar(50), participant_roleReciclator varchar(50), participant_roleEliminator varchar(50), participant_roleAgent varchar(50), participant_roleNegotiator varchar(50), participant_roleBuyer varchar(50), participant_roleSeller varchar(50));
INSERT INTO tokenoffchain (last_block_stored) VALUES(4);