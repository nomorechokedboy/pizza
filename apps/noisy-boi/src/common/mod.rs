use std::{collections::HashMap, hash::Hash};

pub trait FromDBFlatten
where
    Self: From<<Self as FromDBFlatten>::DatabaseType>,
{
    type DatabaseType;
    type GroupBy: Hash + Eq;
    type VecType: From<Self::DatabaseType>;

    fn group_by_field(db_item: &Self::DatabaseType) -> Self::GroupBy;

    fn vec_field(&mut self) -> &mut Vec<Self::VecType>;

    fn group_by(items: Vec<Self::DatabaseType>) -> HashMap<Self::GroupBy, Self> {
        let mut m: HashMap<Self::GroupBy, Self> = HashMap::new();
        for db_item in items {
            match m.entry(Self::group_by_field(&db_item)) {
                std::collections::hash_map::Entry::Occupied(mut o) => {
                    Self::vec_field(o.get_mut()).push(db_item.into())
                }
                std::collections::hash_map::Entry::Vacant(v) => {
                    v.insert(db_item.into());
                }
            }
        }
        m
    }

    fn flatten(items: Vec<Self::DatabaseType>) -> Vec<Self> {
        Self::group_by(items).into_iter().map(|(_, v)| v).collect()
    }

    fn flatten_one(items: Vec<Self::DatabaseType>) -> Option<Self> {
        Self::group_by(items).into_iter().next().map(|(_, v)| v)
    }
}

pub trait Pagination {
    fn pagination() -> () {
        unimplemented!();
    }
}
