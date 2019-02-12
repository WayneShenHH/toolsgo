package dbviews

// SampleView sample
var SampleView = `
DROP VIEW IF EXISTS sample_view;
CREATE VIEW sample_view AS
SELECT 
	*
FROM clock_ins c;`
